<script lang="ts">
  import { setContext } from "svelte";
  import type { NavBarProps, NavBarContext } from "./types";
  import { NAV_BAR_CONTEXT_NAME } from "./consts";

  let { routes, defaultRouteName = routes[0].name }: NavBarProps = $props()
  let context = $state<NavBarContext>({
    selectedRoute: defaultRouteName,
    setSelectedRoute: (newRoute: string) => (context.selectedRoute = newRoute)
  })
  let SelectedComponent = $state(routes.find(route => context.selectedRoute.includes(route.name))?.component);

  $effect(() => {
    SelectedComponent = routes.find(route => context.selectedRoute.includes(route.name))?.component;
  });

  setContext(NAV_BAR_CONTEXT_NAME, context)
</script>

<div class="container">
  <div class="top">
    {#each routes as route}
      {#if route.hasButton}
        <button 
          class={`button ${context.selectedRoute === route.name ? "active" : ""}`} 
          onclick={() => context.setSelectedRoute(route.name)}
        >
          {route.name}
        </button>
      {/if}
    {/each}
  </div>
  <div class="bottom">
    {#if SelectedComponent !== undefined}
      <SelectedComponent />
    {/if}
  </div>
</div>

<style>
  @import url('../../helpers/Constants.css');

  .container {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .top {
    display: flex;
    padding-left: 2px;
    height: 25px;
  }
    
  .top .button {
    all: unset;
    margin: 2px;
    border: 1px solid white;
    background-color: var(--secondary);
    width: 100px;
    height: 25px;
  }

  .top .button:hover {
    background-color: var(--accent);
    cursor: pointer;
  }

  .top .button.active {
    background-color: var(--accent);
  }

  .bottom {
    margin: 0;
    padding: 0;
    margin-top: 10px;
  }
</style>