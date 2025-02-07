<script lang="ts">
  import type { TooltipProps } from "./types";

	let { content, children }: TooltipProps = $props();
	let isHovered = $state(false);
	let x = $state(0);
	let y = $state(0);
	
	function mouseOver(event: MouseEvent) {
		isHovered = true;
		x = event.pageX + 5;
		y = event.pageY + 5;
	}
	function mouseMove(event: MouseEvent) {
		x = event.pageX + 5;
		y = event.pageY + 5;
	}
	function mouseLeave() {
		isHovered = false;
	}
</script>

<div
	role="tooltip"
	onmouseover={mouseOver}
	onfocus={() => {}}
	onmousemove={mouseMove}
	onmouseleave={mouseLeave}
>
	{@render children?.()}
</div>

{#if isHovered}
	<div style="top: {y}px; left: {x}px;" class="tooltip">
        {@render content()}
    </div>
{/if}

<style>
	.tooltip {
		border: 1px solid #ddd;
		box-shadow: 1px 1px 1px #ddd;
		background: white;
		border-radius: 4px;
		padding: 4px;
		position: absolute;
	}
</style>