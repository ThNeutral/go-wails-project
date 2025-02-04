<script>
  import Home from "./components/Home/Home.svelte";
  import NavBar from "./components/NavBar/NavBar.svelte";
  import Weapons from "./components/Weapons/Weapons.svelte";
  
  const routes = [
    {
      name: "Weapons",
      component: Weapons
    }
  ]

</script>

<NavBar routes={routes} />