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

  let countryArray :string[]       // Ex. USA
  let regionArray :string[]       // EX. Texas
  let municipalityArray :string[] // Ex. Bell County
  let municipalityArray2 :string[] // Ex. Temple
  let country :string
  let region :string
  let municipality :string
  let municipality2 :string


  onMount(async function () {
    try {
      const response = await fetch(`http://${window.location.hostname}:8000/Flags/Regional/flag_list.txt`);
      if (response.ok) {
        const text = await response.text();
        countryArray = text.split('\n').slice(0, -1); // This will log an array of strings
      } else {
        console.error('Error fetching text file. Status:', response.status);
      }
    } catch (error) {
      console.error('Error fetching text file:', error);
    }
  });

  const getCountryRegions= async (event) => {
    country = event.target.value
    try {
      const response = await fetch(`http://${window.location.hostname}:8000/Flags/Regional/${country}/flag_list.txt`);
      if (response.ok) {
        const text = await response.text();
        regionArray = text.split('\n').slice(0, -1);
      } else {
        console.error('Error fetching text file. Status:', response.status);
      }
    } catch (error) {
      console.error('Error fetching text file:', error);
    }
  }
  const getMunicipalityFlags = async (event) => {
    region = event.target.value
    try {
      const response = await fetch(`http://${window.location.hostname}:8000/Flags/Regional/${country}/${municipality}/flag_list.txt`);
      if (response.ok) {
        const text = await response.text();
        municipalityArray = text.split('\n').slice(0, -1);
        console.log(municipalityArray);
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
                  <div class="flex items-center">
                </div>
                {#if $userSettings.Theme !== undefined}
                <p>Currently using {$userSettings.Theme.charAt(0).toUpperCase() + $userSettings.Theme.slice(1)}</p>
                {/if}
              </label>


                <label class="label m-6">
                  <h1 class="h2  m-6">Extraflags</h1>
                  <div class="flex space-x-1 justify-center items-center">

                    {#if countryArray !== undefined && countryArray.length>0}
                      <select class="select" on:change={getCountryRegions} style="text-align-last: center; max-width: 200px">
                        {#each countryArray as country}
                          <option value={country}> {country}</option>
                        {/each}
                      </select>

                    {/if}

                  {#if regionArray !== undefined && regionArray.length>0}
                    <select class="select" on:change={null} style="text-align-last: center; max-width: 200px">
                      {#each regionArray as region}
                        <option value="" selected disabled hidden> </option>
                        <option value="\n"> {region}</option>
                      {/each}
                    </select>
                  {/if}

                    <!--{#if regionArray !== undefined && regionArray.length>0}-->
                    <!--  <select class="select" on:change={getMunicipalityFlags} style="text-align-last: center; max-width: 200px">-->
                    <!--    {#each regionArray as region}-->
                    <!--      <option value="\n"> {region}</option>-->
                    <!--    {/each}-->
                    <!--  </select>-->
                    <!--{/if}-->

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