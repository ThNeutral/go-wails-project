import WeaponCard from "../WeaponCard/WeaponCard.svelte";
import WeaponsList from "../WeaponsList/WeaponsList.svelte";
import type { NavBarRoute } from "./types";

export const NAV_BAR_CONTEXT_NAME = "routes";

export const WEAPONS_LIST_ROUTE_NAME = "Weapons";
export const WEAPON_CARD_ROUTE_NAME = "Weapon/";

export const APP_ROUTES: NavBarRoute[] = [
  {
    name: WEAPONS_LIST_ROUTE_NAME,
    component: WeaponsList,
    hasButton: true,
  },
  {
    name: WEAPON_CARD_ROUTE_NAME,
    component: WeaponCard,
  },
] as const;
