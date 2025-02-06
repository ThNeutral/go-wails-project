<script lang="ts">
  import { GetAllWeapons, GetWeaponsByQuery } from "../../../wailsjs/go/main/App";
  import { main } from "../../../wailsjs/go/models";
  import WeaponsMap from "./WeaponsMap.svelte";
  import { debounce } from "../../helpers/Utils";
  import Tooltip from "../Tooltip/Tooltip.svelte";

  let weapons: main.Weapon[] = $state([])
  let isLoading = $state(true)
  let error = $state("")

  let query: string = $state("")
  async function fetchByNameOrID() {
    try {
      const { Data, Error } = await GetWeaponsByQuery(query)
      if (Error && Error.Error) {
        error = Error.Error;
      } else {
        weapons = Data;
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

  const debouncedFetch = debounce(fetchByNameOrID, {
    timeout: 500, 
    before: () => {
      isLoading = true
      error = ""
    }
  })

  fetchByNameOrID()
</script>

<!-- {#snippet triggerContent({ mouseOver })}
  <p>Hover Me</p> -->
<!-- {/snippet} -->

<div class="container">
  <h1>Weapons List</h1>
  <div>
    <input class="input" bind:value={query} onkeyup={debouncedFetch} placeholder="Enter name or ID"/>
    <!-- <Tooltip triggerContent={triggerContent} /> -->
  </div>
  <WeaponsMap isLoading={isLoading} items={weapons} error={error}/>
</div>

<style>
  .input {
    width: 180px;
    padding: 5px 20px;
  }
</style>