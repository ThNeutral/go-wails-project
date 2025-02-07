<script lang="ts">
  import Circle from "svelte-loading-spinners/Circle.svelte";
  import { Colors } from "../../helpers/Constants";
  import type { WeaponsMapProps } from "./types";
  import { getContext } from "svelte";
  import { NAV_BAR_CONTEXT_NAME, WEAPON_CARD_ROUTE_NAME } from "../NavBar/consts";
  import type { NavBarContext } from "../NavBar/types";

  let { isLoading, items, error }: WeaponsMapProps = $props()

  let { selectedRoute, setSelectedRoute } = getContext<NavBarContext>(NAV_BAR_CONTEXT_NAME);
</script>

<div class="container">
  {#if error}
    <div>{error}</div>
  {:else if isLoading}
    <Circle color={Colors.NEUTRAL}/>
  {:else}
    <div class="items">
      {#each items as item}
        <button type="button" class="item" onclick={() => setSelectedRoute(WEAPON_CARD_ROUTE_NAME + item.id)}>
          <img class="image" src={item.assets.image} alt={"weapon image"}/>
          <div>{item.name}</div>
        </button>
      {/each}
    </div>
  {/if}
</div>

<style>
  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 15px;
    margin-bottom: 30px;
  }

  .container .items {
    cursor: pointer;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
  }

  .container .items .item {
    all: unset;
    display: flex;
    flex-direction: column;
    width: 125px;
    height: 125px;
    margin: 12.5px;
  }

  .container .items .item .image {
    width: 100px;
    height: 100px;
  }
</style>