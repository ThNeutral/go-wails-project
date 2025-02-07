<script lang="ts">
  import { GetWeaponTree } from "../../../wailsjs/go/main/App";
  import { main } from "../../../wailsjs/go/models";
  import type { WeaponTreeProps } from "./types";

  const { id }: WeaponTreeProps = $props()

  let treeHead = $state<main.WeaponTreeNode | undefined>()

  async function fetch() {
    const { Data, Error } = await GetWeaponTree(id)
    treeHead = Data
  }
  fetch()
</script>

<div>
  {#if treeHead}
    {JSON.stringify(treeHead, null, 2)}
  {/if}
</div>