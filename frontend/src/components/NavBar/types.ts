import type { Component } from "svelte"

export type NavBarRoute = {
    name: string
    component: Component,
    hasButton?: boolean
}

export type NavBarProps =  {
    routes: NavBarRoute[]
    defaultRouteName?: string
}

export interface NavBarContext {
    selectedRoute: string,
    setSelectedRoute: (newRoute: string) => void
}