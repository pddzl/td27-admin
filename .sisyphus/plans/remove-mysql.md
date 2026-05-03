# Remove MySQL Snippets from AGENTS.md

## TL;DR

> **Quick Summary**: Clean up two MySQL references in AGENTS.md that no longer reflect the project (which uses PostgreSQL exclusively).
>
> **Deliverables**:
> - Updated AGENTS.md with MySQL-free text
>
> **Estimated Effort**: Quick
> **Parallel Execution**: N/A (single task)
> **Critical Path**: One edit task

---

## Context

### Original Request
"Remove all MySQL snippets" from AGENTS.md.

### Details
Two lines reference MySQL:
1. Docker-compose description: `docker-compose/             # MySQL + Redis + Nginx` → should say `PostgreSQL + Redis + Nginx`
2. Gotchas: `- DB defaults to PostgreSQL in config (pgsql section). MySQL config exists but is not actively maintained.` → should only mention PostgreSQL

---

## Work Objectives

### Core Objective
Remove all obsolete MySQL-related text from AGENTS.md.

### Concrete Deliverables
- AGENTS.md: two edits

### Definition of Done
- [ ] grep "MySQL" AGENTS.md returns no matches

### Must Have
- Both MySQL references removed/replaced

### Must NOT Have (Guardrails)
- No other changes to AGENTS.md
- Preserve all PostgreSQL-related information

---

## Verification Strategy

### Test Decision
- **Infrastructure exists**: N/A (markdown doc)
- **Automated tests**: None
- **Agent-Executed QA**: grep for "MySQL" in AGENTS.md

---

## Execution Strategy

Single task, no waves needed.

---

## TODOs

- [ ] 1. Remove MySQL references from AGENTS.md

  **What to do**:
  - Replace `MySQL + Redis + Nginx` with `PostgreSQL + Redis + Nginx` on the docker-compose line
  - Replace the gotcha line `- DB defaults to PostgreSQL in config (pgsql section). MySQL config exists but is not actively maintained.` with `- DB is PostgreSQL (configured under pgsql section in config).`

  **Must NOT do**:
  - Change any other content in AGENTS.md
  - Reformat or restructure the file

  **Recommended Agent Profile**:
  - **Category**: `quick`
    - Reason: Simple two-line text replacement, no logic or testing needed
  - **Skills**: `[]`

  **Parallelization**:
  - **Can Run In Parallel**: N/A
  - **Blocks**: Nothing
  - **Blocked By**: Nothing

  **References**:
  - AGENTS.md:101 — docker-compose line
  - AGENTS.md:106 — gotcha line

  **Acceptance Criteria**:
  - [ ] `grep "MySQL" AGENTS.md` exits with code 1 (no matches found)

  **QA Scenarios**:

  ```
  Scenario: Verify MySQL is gone
    Tool: Bash (grep)
    Preconditions: Edits applied to AGENTS.md
    Steps:
      1. Run: grep -n "MySQL" AGENTS.md
    Expected Result: No output, exit code 1
    Failure Indicators: grep returns any line containing "MySQL"
    Evidence: .sisyphus/evidence/task-1-verify-mysql-removed.txt
  ```

  **Evidence to Capture**:
  - [ ] Run `grep -n "MySQL" AGENTS.md; echo "EXIT:$?"` and capture output

  **Commit**: YES
  - Message: `chore: remove obsolete MySQL references from AGENTS.md`
  - Files: `AGENTS.md`
  - Pre-commit: none

---

## Success Criteria

### Verification Commands
```bash
grep -n "MySQL" AGENTS.md; echo "EXIT:$?"
# Expected: no output, EXIT:1
```

### Final Checklist
- [ ] "MySQL" does not appear in AGENTS.md
- [ ] PostgreSQL info preserved
