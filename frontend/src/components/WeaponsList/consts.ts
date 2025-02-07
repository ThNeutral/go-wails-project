export const WEAPON_SEARCH_FIELDS = [
    { value: "all", label: "All" },
    { value: "name", label: "Name" },
    { value: "id", label: "ID" },
    { value: "slug", label: "Slug" },
    { value: "type", label: "Type" },
    { value: "damagetype", label: "Damage Type" },
    { value: "eldersealtype", label: "Elderseal Type" },
    { value: "rarity", label: "Rarity" }
] as const;

export const WEAPON_INPUT_PLACEHOLDERS: Record<typeof WEAPON_SEARCH_FIELDS[number]["value"], string> = {
    "all": "Searching for all weapons",
    "name": "Enter name of weapon",
    "id": "Enter ID (integer)",
    "slug": "Enter slug",
    "type": "Enter weapon type",
    "damagetype": "Enter damage type",
    "eldersealtype": "Enter elderseal type",
    "rarity": "Enter rarity (integer)"
} as const;