<script lang="ts">
  import { captialize } from "../../helpers/Utils";
  import SharpnessBar from "../Common/SharpnessBar/SharpnessBar.svelte";
  import type { WeaponInfoProps } from "./types";

  const { weapon }: WeaponInfoProps = $props()
</script>

<div class="card">
  <img class="image" alt="weapon" src={weapon.assets.image} />
  <div class="details">
    <p class="text">Type: {weapon.type.split("-").map((part) => captialize(part)).join(" ")}</p>
    <p class="text">Damage: {weapon.attack.display}</p>
    <p class="text">Damage Type: {captialize(weapon.damageType)}</p>
    <p class="text">Affinity: {weapon.attributes.affinity}%</p>
    {#if !weapon?.elements || weapon.elements.length === 0}
      <p class="text">This weapon does not have elemental damage</p>
    {:else}
      {#each weapon.elements as element, index }
        <p class="text">Element {index + 1}: {element.hidden ? "Hidden" : ""} {element.damage} {captialize(element.type)}</p>
      {/each}
    {/if}
    {#if !weapon.elderseal}
      <p class="text">This weapon has no elderseal</p>
    {:else}
      <p class="text">Elderseal: {captialize(weapon.elderseal)}</p>
    {/if}
    {#if !weapon.slots || weapon.slots.length === 0}
      <p class="text">This weapon does not have slots</p>
    {:else}
      <p class="text">Slot{weapon.slots.length !== 1 ? "s" : ""} of {weapon.slots.map((slot) => slot.rank).join(", ")} level{weapon.slots.length !== 1 ? "s" : ""}</p>
    {/if}
    {#if weapon.durability}
      <div class="sharpness">
        <p class="text">Sharpness: </p>
        <SharpnessBar height={16} values={weapon.durability[0]} />
      </div>
    {:else}
      <p class="text">This weapon does not have sharpness</p>
    {/if}
  </div>
</div>

<style>  
  .card {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 16px;
  }

  .details {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
  }

  .text {
    margin: 0;
  }

  .sharpness {
    display: flex;
    gap: 8px;
  }

  .image {
    width: 200px;
    height: 200px;
  }
</style>