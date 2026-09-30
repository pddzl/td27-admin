import type { DictModel } from "@/api/sysManagement/dict"
import type { dictDetailDataModel } from "@/api/sysManagement/dictDetail"
import { dictListApi } from "@/api/sysManagement/dict"
import { dictDetailFlatApi } from "@/api/sysManagement/dictDetail"

export const useDictionaryStore = defineStore("dictionary", () => {
  const dictionaries = ref<DictModel[]>([])
  // cache details by dictId
  const detailsMap = ref<Record<number, dictDetailDataModel[]>>({})

  // In-flight promises: caching only results still lets two concurrent
  // misses fire duplicate requests, so cache the promise itself
  let dictionariesPromise: Promise<void> | null = null
  const detailPromises = new Map<number, Promise<void>>()

  const fetchDictionaries = () => {
    if (dictionaries.value.length > 0) return Promise.resolve()
    if (!dictionariesPromise) {
      dictionariesPromise = dictListApi({})
        .then((res) => {
          if (res.code === 0) {
            dictionaries.value = res.data.list
          }
        })
        .finally(() => {
          dictionariesPromise = null // allow retry after failure
        })
    }
    return dictionariesPromise
  }

  const fetchDictionaryDetail = (dictId: number) => {
    if (detailsMap.value[dictId]) return Promise.resolve() // ✅ cached
    let promise = detailPromises.get(dictId)
    if (!promise) {
      promise = dictDetailFlatApi({ dictId })
        .then((res) => {
          if (res.code === 0) {
            detailsMap.value[dictId] = res.data
          }
        })
        .finally(() => {
          detailPromises.delete(dictId) // allow retry after failure
        })
      detailPromises.set(dictId, promise)
    }
    return promise
  }

  // ✅ Helper: get options by en_name
  const getOptions = async (en_name: string) => {
    // find dictId
    if (dictionaries.value.length === 0) {
      await fetchDictionaries()
    }

    const dict = dictionaries.value.find(d => d.en_name === en_name)
    if (!dict) return []

    // fetch details if needed
    await fetchDictionaryDetail(dict.id)

    return (detailsMap.value[dict.id] || []).map((item: dictDetailDataModel) => ({
      label: item.label,
      value: item.value
    }))
  }

  return {
    dictionaries,
    detailsMap,
    fetchDictionaries,
    fetchDictionaryDetail,
    getOptions // ✅ expose helper
  }
})
