<script lang="ts">
  import { getContext } from "svelte";
  import { NAV_BAR_CONTEXT_NAME, WEAPON_CARD_ROUTE_NAME, WEAPONS_LIST_ROUTE_NAME } from "../NavBar/consts";
  import type { NavBarContext } from "../NavBar/types";
  import { GetWeaponsByQuery } from "../../../wailsjs/go/main/App";
  import { main } from "../../../wailsjs/go/models";
  import { captialize, isErrorObject } from "../../helpers/Utils";
  import { Circle } from "svelte-loading-spinners";
  import { Colors } from "../../helpers/Constants";
  import { Log } from "../../../wailsjs/go/main/Logger";
  import SharpnessBar from "../Common/SharpnessBar/SharpnessBar.svelte";
  import WeaponInfo from "./WeaponInfo.svelte";
  import WeaponTree from "./WeaponTree.svelte";

  let { selectedRoute } = getContext<NavBarContext>(NAV_BAR_CONTEXT_NAME);

  let isLoading = $state(true)
  let error = $state("")
  let weapon = $state<main.Weapon>()

  async function fetch() {
    try {
      Log(selectedRoute.replace(WEAPON_CARD_ROUTE_NAME, ""))
      const { Data, Error } = await GetWeaponsByQuery({ Field: "id", Query: selectedRoute.replace(WEAPON_CARD_ROUTE_NAME, "")})
      if (isErrorObject(Error)) {
        error = Error.Error
      } else {
        weapon = Data[0]
      }
    } catch (e) {
      if (e instanceof Error) {
        error = e.message;
      } else {
        error = "Unknown error happened"
      }
    } finally {
      isLoading = false;
    }
  }
  fetch()
</script>

<div class="container">
  {#if error}
    <div>{error}</div>
  {:else if isLoading}
    <Circle color={Colors.NEUTRAL}/>
  {:else}
    <h1>{weapon!.name}</h1>
    <div class="info">
      <WeaponInfo weapon={weapon!}/>
      <WeaponTree id={weapon!.id}/>
    </div>
  {/if}
</div>

<style>
  .container {
    padding: 20px 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }

  .info {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
  }
</style>