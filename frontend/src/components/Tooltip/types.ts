import type { Snippet } from "svelte";

export interface TooltipProps {
    triggerContent: Snippet<[{mouseOver: (e: MouseEvent) => void}]>
    tooltipContent: Snippet
}