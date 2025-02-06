<script lang="ts">
  import Circle from "svelte-loading-spinners/Circle.svelte";
  import { Colors } from "../../helpers/Constants";
  import type { WeaponsMapProps } from "./types";

  let { isLoading, items, error }: WeaponsMapProps = $props()
</script>

<div class="container">
  {#if error}
    <div>{error}</div>
  {:else if isLoading}
    <Circle color={Colors.NEUTRAL}/>
  {:else}
    <div class="items">
      {#each items as item}
        <div class="item">
          <img class="image" src={item.assets.image} alt={"weapon image"}/>
          <div>{item.name}</div>
        </div>
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
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
  }

  .container .items .item {
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