import type { Component } from "svelte"

export type NavBarRoute = {
    name: string
    component: Component
}

export type NavBarProps =  {
    routes: NavBarRoute[]
}