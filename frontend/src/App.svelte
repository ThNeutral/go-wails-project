<script lang="ts">
  import { APP_ROUTES, WEAPONS_LIST_ROUTE_NAME } from "./components/NavBar/consts";
  import NavBar from "./components/NavBar/NavBar.svelte";
</script>

<NavBar routes={APP_ROUTES} defaultRouteName={WEAPONS_LIST_ROUTE_NAME} />