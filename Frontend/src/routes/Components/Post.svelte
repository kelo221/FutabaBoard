<script lang="ts">
    import ExtraFlags from "./ExtraFlags.svelte";
    export let postType; // Pass 'thread' or 'post' as a prop
    export let content;
    import Icon from "@iconify/svelte";

    export let threadID
    export let isOpen


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

    const handleImageClick = imageHash => {
        enlargedImage = enlargedImage === imageHash ? null : imageHash;
    };



    import { replyBoxStore } from '../../stores'

    const openReply = (threadID :number, postType :string, targetPost :string) => {
        $replyBoxStore.threadID = threadID
        $replyBoxStore.open = true

        if (postType === "post") {
            $replyBoxStore.content =  $replyBoxStore.content + ">>" +targetPost+ "\n"
        }

    }


</script>

<div class="card" style="display: flex; flex-direction: column;">
    <div class="container flex flex-col items-center justify-center m-2">

    {#if postType === 'thread'}
        <h1 style="font-size: x-large">{content.Topic}</h1>
    {/if}
    </div>
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
                        <div class="flex justify-center">
                            <b style={'color: #' + content.UserHash}>{content.Name}</b>
                        </div>
                        <div class="flex space-x-1 justify-center items-center">
                            <i class={'flag ' + content.Country}></i>
                            <ExtraFlags Flags={content.ExtraFlags}/>
                        </div>
                    </div>

                <div class="grid">
                    <p>#{content.ID}</p>
                    <p>{content.PostCount} Posts</p>
                </div>
                    <div class="grid">
                        <p>{parseDateStringToDate(content.UnixTime)}</p>
                        <p style={'color: #' + content.UserHash} >
                            {"#" + content.UserHash} </p>

                    </div>
                {:else}
                    <div class="grid">
                        <b style={'color: #' + content.UserHash}>{content.Name}</b>
                        <div class="flex space-x-1 justify-center items-center">
                            <i class={'flag ' + content.Country}></i>
                            <ExtraFlags Flags={content.ExtraFlags}/>
                        </div>
                    </div>
                    <p>#{content.ID}</p>
                    <p>{parseDateStringToDate(content.UnixTime)}</p>
                    <p style={'color: #' + content.UserHash}>
                        {"#" + content.UserHash} </p>
                {/if}

            {:else}

                <div class="grid">
                    <div class="flex justify-center">
                        <b style={'color: #' + content.UserHash}>{content.Name}</b>
                    </div>
                    <div class="flex space-x-1 justify-center items-center">
                        <i class={'flag ' + content.Country}></i>
                        <ExtraFlags Flags={content.ExtraFlags}/>
                    </div>
                </div>

                <p style="color: blue; text-decoration: underline; cursor: pointer;"  on:click={() => openReply(threadID, postType, content.ID) }>#{content.ID}</p>
            <p>{postType === 'thread' ? content.PostCount + ' Posts' : ''}</p>
            <p>{parseDateStringToDate(content.UnixTime)}</p>
            <p style={'color: #' + content.UserHash}>
                {"#" + content.UserHash}
            </p>

            {/if}

        </div>
        {#if postType === 'thread' && isOpen === false}
            {#if isMobile}
                <button class="btn variant-filled mobile-icon-btn" type="button">
                    <Icon icon="ooui:new-window-rtl" />
                </button>
            {:else}
                <a href={`thread/${content.ID}`}>
                <button class="btn variant-filled" type="button">
                    Open Thread
                </button>
                </a>
            {/if}
        {/if}
    </footer>
</div>
