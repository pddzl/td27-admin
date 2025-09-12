import type { dictDataModel } from "@/api/sysSet/dict"
import type { dictDetailDataModel } from "@/api/sysSet/dictDetail"
import { getDictApi } from "@/api/sysSet/dict"
import { getDictDetailApi } from "@/api/sysSet/dictDetail"

export const useDictionaryStore = defineStore("dictionary", () => {
  const dictionaries = ref<dictDataModel[]>([])
  // cache details by dictId
  const detailsMap = ref<Record<number, dictDetailDataModel[]>>({})

  const fetchDictionaries = async () => {
    if (dictionaries.value.length > 0) return
    try {
      const res = await getDictApi()
      if (res.code === 0) {
        dictionaries.value = res.data
      }
    } finally {
      //
    }
  }

  const fetchDictionaryDetail = async (dictId: number) => {
    if (detailsMap.value[dictId]) return // ✅ cached
    try {
      const res = await getDictDetailApi({
        page: 0,
        pageSize: 0,
        dictId
      })
      if (res.code === 0) {
        detailsMap.value[dictId] = res.data.list
      }
    } finally {
      //
    }
  }

  // ✅ Helper: get options by enName
  const getOptions = async (enName: string) => {
    // find dictId
    if (dictionaries.value.length === 0) {
      await fetchDictionaries()
    }
    const dict = dictionaries.value.find(d => d.enName === enName)
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
