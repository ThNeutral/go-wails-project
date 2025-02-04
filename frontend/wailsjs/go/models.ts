export namespace main {
	
	export class Attack {
	    display: number;
	    raw: number;
	
	    static createFrom(source: any = {}) {
	        return new Attack(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.display = source["display"];
	        this.raw = source["raw"];
	    }
	}
	export class Item {
	    id: number;
	    name: string;
	    description: string;
	    rarity: number;
	    carryLimit: number;
	    sellPrice: number;
	    buyPrice: number;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.rarity = source["rarity"];
	        this.carryLimit = source["carryLimit"];
	        this.sellPrice = source["sellPrice"];
	        this.buyPrice = source["buyPrice"];
	    }
	}
	export class CraftingCost {
	    quantity: number;
	    item: Item;
	
	    static createFrom(source: any = {}) {
	        return new CraftingCost(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.quantity = source["quantity"];
	        this.item = this.convertValues(source["item"], Item);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class Slot {
	    rank: number;
	
	    static createFrom(source: any = {}) {
	        return new Slot(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rank = source["rank"];
	    }
	}
	export class WeaponAttributes {
	    affinity: number;
	    defense: number;
	
	    static createFrom(source: any = {}) {
	        return new WeaponAttributes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.affinity = source["affinity"];
	        this.defense = source["defense"];
	    }
	}
	export class WeaponSharpness {
	    red: number;
	    orange: number;
	    yellow: number;
	    green: number;
	    blue: number;
	    white: number;
	    purple: number;
	
	    static createFrom(source: any = {}) {
	        return new WeaponSharpness(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.red = source["red"];
	        this.orange = source["orange"];
	        this.yellow = source["yellow"];
	        this.green = source["green"];
	        this.blue = source["blue"];
	        this.white = source["white"];
	        this.purple = source["purple"];
	    }
	}
	export class WeaponAssets {
	    icon: string;
	    image: string;
	
	    static createFrom(source: any = {}) {
	        return new WeaponAssets(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.icon = source["icon"];
	        this.image = source["image"];
	    }
	}
	export class WeaponCraftingInfo {
	    craftable: boolean;
	    previous: number;
	    branches: number[];
	    craftingMaterials: CraftingCost[];
	    upgradeMaterials: CraftingCost[];
	
	    static createFrom(source: any = {}) {
	        return new WeaponCraftingInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.craftable = source["craftable"];
	        this.previous = source["previous"];
	        this.branches = source["branches"];
	        this.craftingMaterials = this.convertValues(source["craftingMaterials"], CraftingCost);
	        this.upgradeMaterials = this.convertValues(source["upgradeMaterials"], CraftingCost);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class WeaponElement {
	    type: string;
	    damage: number;
	    hidden: boolean;
	
	    static createFrom(source: any = {}) {
	        return new WeaponElement(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.damage = source["damage"];
	        this.hidden = source["hidden"];
	    }
	}
	export class Weapon {
	    id: number;
	    slug: string;
	    name: string;
	    type: string;
	    rarity: number;
	    attack: Attack;
	    slots: Slot[];
	    elements: WeaponElement[];
	    crafting: WeaponCraftingInfo;
	    assets: WeaponAssets;
	    durability: WeaponSharpness[];
	    elderseal: string;
	    damageType: string;
	    attributes: WeaponAttributes;
	
	    static createFrom(source: any = {}) {
	        return new Weapon(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.slug = source["slug"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.rarity = source["rarity"];
	        this.attack = this.convertValues(source["attack"], Attack);
	        this.slots = this.convertValues(source["slots"], Slot);
	        this.elements = this.convertValues(source["elements"], WeaponElement);
	        this.crafting = this.convertValues(source["crafting"], WeaponCraftingInfo);
	        this.assets = this.convertValues(source["assets"], WeaponAssets);
	        this.durability = this.convertValues(source["durability"], WeaponSharpness);
	        this.elderseal = source["elderseal"];
	        this.damageType = source["damageType"];
	        this.attributes = this.convertValues(source["attributes"], WeaponAttributes);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	

}

