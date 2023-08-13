<script lang="ts">



  import {localStorageStore} from "@skeletonlabs/skeleton"; //Backup data
  import type { Writable } from "svelte/store";
  import type { UserPrefences } from "../../app";

  const userSettings : Writable<UserPrefences> = localStorageStore("userPrefs", {});
  function changeTheme(event) {
    const selectedTheme = event.target.value;
    console.log(selectedTheme);
    $userSettings.Theme = event.target.value
    document.getElementById("themeHooker").setAttribute('data-theme', selectedTheme);
  }


</script>

<main class="m-3">
  <div class="container">
    <div class="flex-1 space-y-3">
      <div class="container h-full mx-auto mt-4 flex justify-center items-center">
        <div class="space-y-5 text-center flex flex-col items-center">
          <div class="card">

            <div>
              <h1 class="h1 m-6">Settings</h1>
              <hr class="full-width-hr" />

              <label class="label m-6">
                <h1 class="h2  m-6">Theme</h1>
                <select class="select" on:change={changeTheme} style="text-align-last: center">
                  <option value="skeleton">Skeleton</option>
                  <option value="modern">Modern</option>
                  <option value="rocket">Rocket</option>
                  <option value="seafoam">Seafoam</option>
                  <option value="vintage">Vintage</option>
                  <option value="sahara">Sahara</option>
                  <option value="hamlindigo">Hamlindigo</option>
                  <option value="gold-nouveau">Gold Nouveau</option>
                  <option value="crimson">Crimson</option>
                </select>
              </label>
            </div>

            <label class="label m-6">
              <h1 class="h2  m-6">Image Rights Key</h1>
              <input class="input" placeholder="" style="text-align: center" type="text">
            </label>

            <label class="label m-6">
              <h1 class="h2 m-6">Filename preference</h1>
              <div class="space-y-2">
                <label class="flex items-center space-x-2">
                  <input checked class="radio" name="radio-direct" type="radio" value="1" />
                  <p>Original filename</p>
                </label>
                <label class="flex items-center space-x-2">
                  <input class="radio" name="radio-direct" type="radio" value="2" />
                  <p>Prefer hash as filename</p>
                </label>
              </div>
            </label>

          </div>
        </div>
      </div>
    </div>
  </div>
</main>
