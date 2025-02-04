<script lang="ts">
  import { GetWeaponsByName } from "../../../wailsjs/go/main/App";
  import { GetAllWeapons, GetWeaponByID } from "../../../wailsjs/go/main/App";
  import { main } from "../../../wailsjs/go/models";
  import WeaponsMap from "./WeaponsMap.svelte";
  import { debounce } from "../../helpers/Utils";

  let weapons: main.Weapon[] = $state([])
  let isLoading = $state(true)
  let error = $state<Error>()

  let query: string = $state("")
  async function fetchByNameOrID() {
    try {
      if (query === "") {
        weapons = await GetAllWeapons()
      } else if (!Number.isNaN(+query)) {
        weapons = [await GetWeaponByID(+query)]
      } else {
        weapons = await GetWeaponsByName(query)
      }
    } catch(e) {
      error = e as Error
    } finally {
      isLoading = false
    }
  }

  const debouncedFetch = debounce(fetchByNameOrID, {
    timeout: 500, 
    before: () => {
      isLoading = true
      error = undefined
    }
  })

  fetchByNameOrID()
</script>

<div class="container">
  <h1>Weapons List</h1>
  <input bind:value={query} onkeyup={debouncedFetch} placeholder="Enter name or ID"/>
  <WeaponsMap isLoading={isLoading} items={weapons}/>
</div>

<style>
</style>