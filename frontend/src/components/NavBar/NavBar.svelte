<script lang="ts">
  import type { NavBarProps } from "./types";

  let { routes }: NavBarProps = $props()
  let selectedIndex = $state(0)
  const SelectedComponent = $derived(routes[selectedIndex].component)
</script>

<div class="container">
  <div class="top">
    {#each routes as route, index}
      <button 
        class={`button ${selectedIndex === index ? "active" : ""}`} 
        onclick={() => selectedIndex = index}
      >
        {route.name}
      </button>
    {/each}
  </div>
  <div class="bottom">
    <SelectedComponent />
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