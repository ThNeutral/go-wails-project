import { main } from "../../../wailsjs/go/models"

export type WeaponsMapProps = {
    isLoading: boolean
    items: main.Weapon[]
    error?: Error
}