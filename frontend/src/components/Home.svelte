<script lang="ts">
    import { main } from '../../wailsjs/go/models.js';
    import { GetWeaponByID } from '../../wailsjs/go/main/App.js'
  
    let weapon: main.Weapon
    let id: string
  
    function greet() {
      if (id) {
        GetWeaponByID(id).then(result => weapon = result)
      }
    }
  
    function prettyPring(obj?: object) {
      if (!obj) return ""
      return JSON.stringify(obj, null, 2)
    }
  </script>
  
  <main>
    <div class="input-box" id="input">
      <input autocomplete="off" bind:value={id} class="input" id="name" type="text"/>
      <button class="btn" on:click={greet}>Greet</button>
    </div>
    <div class="result" id="result">{prettyPring(weapon)}</div>
  </main>
  
  <style>
  
    #logo {
      display: block;
      width: 50%;
      height: 50%;
      margin: auto;
      padding: 10% 0 0;
      background-position: center;
      background-repeat: no-repeat;
      background-size: 100% 100%;
      background-origin: content-box;
    }
  
    .result {
      height: 20px;
      line-height: 20px;
      margin: 1.5rem auto;
    }
  
    .input-box .btn {
      width: 60px;
      height: 30px;
      line-height: 30px;
      border-radius: 3px;
      border: none;
      margin: 0 0 0 20px;
      padding: 0 8px;
      cursor: pointer;
    }
  
    .input-box .btn:hover {
      background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
      color: #333333;
    }
  
    .input-box .input {
      border: none;
      border-radius: 3px;
      outline: none;
      height: 30px;
      line-height: 30px;
      padding: 0 10px;
      background-color: rgba(240, 240, 240, 1);
      -webkit-font-smoothing: antialiased;
    }
  
    .input-box .input:hover {
      border: none;
      background-color: rgba(255, 255, 255, 1);
    }
  
    .input-box .input:focus {
      border: none;
      background-color: rgba(255, 255, 255, 1);
    }
  
  </style>
  