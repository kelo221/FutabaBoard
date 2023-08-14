<script lang="ts">
  import { localStorageStore } from "@skeletonlabs/skeleton"; //Backup data
  import type { Writable } from "svelte/store";
  import type { UserPrefences } from "../../app";
  import { onMount } from "svelte";

  const userSettings : Writable<UserPrefences> = localStorageStore("userPrefs", {});
  function changeTheme(event) {
    const selectedTheme = event.target.value;
    console.log(selectedTheme);
    $userSettings.Theme = event.target.value
    document.getElementById("themeHooker").setAttribute('data-theme', selectedTheme);
  }

  let countryList :string[]
  let state :string[]
  let regionArray :string[]
  let region02Array :string[]
  let Country = ""

  onMount(async function () {
    try {
      const response = await fetch(`http://${window.location.hostname}:8000/Flags/Regional/flag_list.txt`);
      if (response.ok) {
        const text = await response.text();
        countryList = text.split('\n').slice(0, -1); // This will log an array of strings
      } else {
        console.error('Error fetching text file. Status:', response.status);
      }
    } catch (error) {
      console.error('Error fetching text file:', error);
    }
  });


  const getAdditionFlags= async (event) => {
    Country = event.target.value + "/";
    regionArray = undefined
    
    try {
      const response = await fetch(`http://${window.location.hostname}:8000/Flags/Regional/${Country}flag_list.txt`);
      if (response.ok) {
        const text = await response.text();
        state = text.split('\n').slice(0, -1);
      } else {
        console.error('Error fetching text file. Status:', response.status);
      }
    } catch (error) {
      console.error('Error fetching text file:', error);
    }
  }

</script>
<main class="m-3">
  <div class="container h-full mx-auto mt-4 flex justify-center items-center">
    <div class="space-y-5 text-center flex flex-col items-center">
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


                <label class="label m-6">
                  <h1 class="h2  m-6">Extraflags</h1>
                  <div class="flex space-x-1 justify-center items-center">

                    {#if countryList !== undefined && countryList.length>0}

                      <select class="select" on:change={getAdditionFlags} style="text-align-last: center; max-width: 200px">
                        {#each countryList as country}
                          <option value={country}> {country}</option>
                        {/each}
                      </select>

                    {/if}

                  {#if state !== undefined && state.length>0}
                    <select class="select" on:change={undefined} style="text-align-last: center; max-width: 200px">
                      {#each state as localState}
                        <option value="" selected disabled hidden> </option>
                        <option value="\n"> {localState}</option>
                      {/each}
                    </select>
                  {/if}

                    {#if regionArray !== undefined && regionArray.length>0}
                      <select class="select" on:change={undefined} style="text-align-last: center; max-width: 200px">
                        {#each regionArray as region}
                          <option value="\n"> {region}</option>
                        {/each}
                      </select>
                    {/if}

                    {#if region02Array !== undefined}
                      <select class="select"  style="text-align-last: center; max-width: 200px">
                        {#each region02Array as region}
                          <option value="\n"> {region}</option>
                        {/each}
                      </select>
                    {/if}
                </label>

              <button type="button" class="btn btn-sm variant-filled">Set Flags</button>
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