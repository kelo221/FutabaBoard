<script lang="ts">
    export let postType; // Pass 'thread' or 'post' as a prop
    export let content;
    import Icon from "@iconify/svelte";

    export let threadID

    let isMobile = false;

    if (window.innerWidth <= 768) {
        isMobile = true;
    }

    function parseDateStringToDate(datetimeString: string): string {
        const date = new Date(datetimeString);
        const options: Intl.DateTimeFormatOptions = {
            year: '2-digit',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false
        };

        return (
            date.toLocaleDateString('en-US', options).replace(',', '')
        );
    }

    let enlargedImage = null;

    function handleImageClick(imageHash) {
        enlargedImage = enlargedImage === imageHash ? null : imageHash;
    }

</script>

<div class="card" style="display: flex; flex-direction: column;">

    <div class="flex" style="text-align: left; flex-grow: 1;">

        {#if !isMobile}
            <div class="container m-4" style="text-align: left; flex-grow: 1;">
                {@html content.Text}
            </div>
        {:else}
            {#if !enlargedImage}
            <div class="container m-4" style="text-align: left; flex-grow: 1;">
                {@html content.Text}
            </div>
            {/if}
        {/if}


            {#if content.PostImage.ImageHash }
                <div class="container flex flex-col items-center justify-center m-4">
                    <img
                            src={`http://localhost:8000/ThreadContent/${threadID}/${content.PostImage.ImageHash}${
                                enlargedImage === content.PostImage.ImageHash ? content.PostImage.ImageInfo.split(" ").pop() : '_small.png'}`}
                            alt="Post Image"

                            on:click={() => handleImageClick(content.PostImage.ImageHash)}
                    />
                    {#if !enlargedImage}
                    <small style="text-align: center">{content.PostImage.ImageInfo}</small>
                    {/if}
                </div>
            {/if}
    </div>

    <hr class="opacity-50"/>
    <footer class="p-4 flex justify-start items-center space-x-4" style="margin-top: auto;">
        <div class="flex-auto flex justify-between items-center">

            {#if isMobile}

                {#if postType === 'thread'}
                <div class="grid">
                    <b style={'color: #' + content.Hash}>{content.Name}</b>
                    <div class="flex space-x-1">
                        <i class={'flag ' + content.Country}></i>
                        <img src={`http://${window.location.hostname}:8000/Flags/Regional/${content.ExtraFlags}.png`}/>
                    </div>
                </div>

                <div class="grid">
                    <p>#{content.ID}</p>
                    <p>{content.PostCount} Posts</p>

                </div>

                    <div class="grid">
                        <p>{parseDateStringToDate(content.UnixTime)}</p>
                        <p style={'color: #' + content.UserHash}>
                            {"#" + content.UserHash} </p>
                    </div>



                {:else}

                    <div class="grid">
                        <b style={'color: #' + content.Hash}>{content.Name}</b>
                        <div class="flex space-x-1">
                            <i class={'flag ' + content.Country}></i>
                            <img src={`http://${window.location.hostname}:8000/Flags/Regional/${content.ExtraFlags}.png`}/>
                        </div>
                    </div>

                    <p>#{content.ID}</p>
                    <p>{parseDateStringToDate(content.UnixTime)}</p>
                    <p style={'color: #' + content.UserHash}>
                        {"#" + content.UserHash} </p>
                {/if}

            {:else}

            <div class="grid">
                <b style={'color: #' + content.Hash}>{content.Name}</b>
                <div class="flex space-x-1">
                    <i class={'flag ' + content.Country}></i>
                    <img src={`http://${window.location.hostname}:8000/Flags/Regional/${content.ExtraFlags}.png`}/>
                </div>
            </div>
            <p>#{content.ID}</p>
            <p>{postType === 'thread' ? content.PostCount + ' Posts' : ''}</p>
            <p>{parseDateStringToDate(content.UnixTime)}</p>
            <p style={'color: #' + content.UserHash}>
                {"#" + content.UserHash}
            </p>

            {/if}

        </div>
        {#if postType === 'thread'}
            {#if isMobile}
                <button class="btn variant-filled mobile-icon-btn" type="button">
                    <Icon icon="ooui:new-window-rtl" />
                </button>
            {:else}
                <button class="btn variant-filled" type="button">
                    Open Thread
                </button>
            {/if}
        {/if}
    </footer>
</div>
