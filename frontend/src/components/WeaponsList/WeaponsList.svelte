<script lang="ts">
  import { GetAllWeapons, GetWeaponsByQuery } from "../../../wailsjs/go/main/App";
  import { main } from "../../../wailsjs/go/models";
  import WeaponsMap from "./WeaponsMap.svelte";
  import { debounce, isErrorObject } from "../../helpers/Utils";
  import Tooltip from "../Tooltip/Tooltip.svelte";
  import { WEAPON_INPUT_PLACEHOLDERS, WEAPON_SEARCH_FIELDS } from "./consts";

  let weapons: main.Weapon[] = $state([])
  let isLoading = $state(true)
  let error = $state("")

  let field: typeof WEAPON_SEARCH_FIELDS[number]["value"] = $state("all")
  let query = $state("")

  $effect(() => {
    if (field === "all") {
      query = ""
      fetchByQuery()
    }
  })

  async function fetchByQuery() {
    try {
      let data: main.Weapon[], err: any;
      if (field == "all") {
        const { Data, Error } = await GetAllWeapons()
        data = Data; err = Error;
      } else {
        const { Data, Error } = await GetWeaponsByQuery({ Field: field, Query: query })
        data = Data; err = Error;  
      }
      if (isErrorObject(err)) {
        error = err.Error;
      } else {
        weapons = data;
      }
    } catch(e) {
      if (e instanceof Error) {
        error = e.message;
      } else {
        error = "Unknown error happened"
      }
    } finally {
      isLoading = false
    }
  }

  const debouncedFetch = debounce(fetchByQuery, {
    timeout: 500, 
    before: () => {
      isLoading = true
      error = ""
    }
  })

  fetchByQuery()
</script>

<div class="container">
  <h1>Weapons List</h1>
  <div>
    <select class="dropdown" bind:value={field}>
      {#each WEAPON_SEARCH_FIELDS as field, index }
        <option value={field.value} selected={index === 0}>{field.label}</option>
      {/each}
    </select>
    <input disabled={field === "all"} class="input" bind:value={query} onkeyup={debouncedFetch} placeholder={WEAPON_INPUT_PLACEHOLDERS[field]}/>
  </div>
  <WeaponsMap isLoading={isLoading} items={weapons} error={error}/>
</div>

<style>
  .dropdown, .input {
    padding: 5px 20px;
  }
  .input {
    width: 180px;
  }
</style>