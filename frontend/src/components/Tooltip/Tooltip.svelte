<script lang="ts">
  import { TooltipProps } from "./types";

	let { tooltipContent: TooltipContent, triggerContent: TriggerContent }: TooltipProps = $props();
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

{@render TooltipContent({ mouseOver })}

{#if isHovered}
	<div style="top: {y}px; left: {x}px;" class="tooltip">
        <TooltipContent />
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