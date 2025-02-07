<script lang="ts">
  import { Log } from "../../../../wailsjs/go/main/Logger";
  import { main } from "../../../../wailsjs/go/models";
  import type { SharpnessBarProps } from "./types";

  const { values, height = 12, width = 100 }: SharpnessBarProps = $props()

  function getSharpnessInPercentages(value: number) {
    return Math.floor(value / 400 * 100);
  }

  function getStyles(key: keyof main.WeaponSharpness) {
    const percentage = getSharpnessInPercentages(values[key as keyof main.WeaponSharpness]);
    return (`
      height:${height}px;
      width:${percentage * width / 100}px;
      background-color:${key.toLowerCase()};
    `);
  }

  function getLeftoverStyles() {
    const leftoverValue = Object.keys(values).reduce((prev, next) => (prev - values[next as keyof main.WeaponSharpness]), 400)
    const percentage = getSharpnessInPercentages(leftoverValue)
    return (`
      height:${height}px;
      width:${percentage * width / 100}px;
      background-color:black;
    `);
  }
</script>

<div class="bar" style={`height:${height}px;width:${width};`}>
  {#each Object.keys(values) as key}
    <div style={getStyles(key as keyof main.WeaponSharpness)}></div>
  {/each}
  <div style={getLeftoverStyles()}></div>
</div>

<style>
  .bar {
    display: flex;
    background-color: black;
    border: 3px solid black;
  }
</style>