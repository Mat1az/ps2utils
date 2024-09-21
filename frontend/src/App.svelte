<script>
    import {RepairFile, SelectFile} from '../wailsjs/go/main/App';


    let games = []
    let index = 1

    let items = [
        {
            id: 0,
            name: "PS2 Manager",
            v: "M5 11.424V1a1 1 0 1 0-2 0v10.424a3.228 3.228 0 0 0 0 6.152V19a1 1 0 1 0 2 0v-1.424a3.228 3.228 0 0 0 0-6.152ZM19.25 14.5A3.243 3.243 0 0 0 17 11.424V1a1 1 0 0 0-2 0v10.424a3.227 3.227 0 0 0 0 6.152V19a1 1 0 1 0 2 0v-1.424a3.243 3.243 0 0 0 2.25-3.076Zm-6-9A3.243 3.243 0 0 0 11 2.424V1a1 1 0 0 0-2 0v1.424a3.228 3.228 0 0 0 0 6.152V19a1 1 0 1 0 2 0V8.576A3.243 3.243 0 0 0 13.25 5.5Z"
        },
        {
            id: 1,
            name: "Games Manager",
            v: "M6.143 0H1.857A1.857 1.857 0 0 0 0 1.857v4.286C0 7.169.831 8 1.857 8h4.286A1.857 1.857 0 0 0 8 6.143V1.857A1.857 1.857 0 0 0 6.143 0Zm10 0h-4.286A1.857 1.857 0 0 0 10 1.857v4.286C10 7.169 10.831 8 11.857 8h4.286A1.857 1.857 0 0 0 18 6.143V1.857A1.857 1.857 0 0 0 16.143 0Zm-10 10H1.857A1.857 1.857 0 0 0 0 11.857v4.286C0 17.169.831 18 1.857 18h4.286A1.857 1.857 0 0 0 8 16.143v-4.286A1.857 1.857 0 0 0 6.143 10Zm10 0h-4.286A1.857 1.857 0 0 0 10 11.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 18 16.143v-4.286A1.857 1.857 0 0 0 16.143 10Z"
        },
        {
            id: 2,
            name: "POPS Manager",
            v: "M6.143 0H1.857A1.857 1.857 0 0 0 0 1.857v4.286C0 7.169.831 8 1.857 8h4.286A1.857 1.857 0 0 0 8 6.143V1.857A1.857 1.857 0 0 0 6.143 0Zm10 0h-4.286A1.857 1.857 0 0 0 10 1.857v4.286C10 7.169 10.831 8 11.857 8h4.286A1.857 1.857 0 0 0 18 6.143V1.857A1.857 1.857 0 0 0 16.143 0Zm-10 10H1.857A1.857 1.857 0 0 0 0 11.857v4.286C0 17.169.831 18 1.857 18h4.286A1.857 1.857 0 0 0 8 16.143v-4.286A1.857 1.857 0 0 0 6.143 10Zm10 0h-4.286A1.857 1.857 0 0 0 10 11.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 18 16.143v-4.286A1.857 1.857 0 0 0 16.143 10Z"
        },
        {
            id: 3,
            name: "OPL Manager",
            v: "M5 11.424V1a1 1 0 1 0-2 0v10.424a3.228 3.228 0 0 0 0 6.152V19a1 1 0 1 0 2 0v-1.424a3.228 3.228 0 0 0 0-6.152ZM19.25 14.5A3.243 3.243 0 0 0 17 11.424V1a1 1 0 0 0-2 0v10.424a3.227 3.227 0 0 0 0 6.152V19a1 1 0 1 0 2 0v-1.424a3.243 3.243 0 0 0 2.25-3.076Zm-6-9A3.243 3.243 0 0 0 11 2.424V1a1 1 0 0 0-2 0v1.424a3.228 3.228 0 0 0 0 6.152V19a1 1 0 1 0 2 0V8.576A3.243 3.243 0 0 0 13.25 5.5Z"
        },
        {
            id: 4,
            name: "SaveData Manager",
            v: "M16 1h-3.278A1.992 1.992 0 0 0 11 0H7a1.993 1.993 0 0 0-1.722 1H2a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2Zm-3 14H5a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2Zm0-4H5a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2Zm0-5H5a1 1 0 0 1 0-2h2V2h4v2h2a1 1 0 1 1 0 2Z"
        },
        {
            id: 5,
            name: "Device Manager",
            v: "M5 11.424V1a1 1 0 1 0-2 0v10.424a3.228 3.228 0 0 0 0 6.152V19a1 1 0 1 0 2 0v-1.424a3.228 3.228 0 0 0 0-6.152ZM19.25 14.5A3.243 3.243 0 0 0 17 11.424V1a1 1 0 0 0-2 0v10.424a3.227 3.227 0 0 0 0 6.152V19a1 1 0 1 0 2 0v-1.424a3.243 3.243 0 0 0 2.25-3.076Zm-6-9A3.243 3.243 0 0 0 11 2.424V1a1 1 0 0 0-2 0v1.424a3.228 3.228 0 0 0 0 6.152V19a1 1 0 1 0 2 0V8.576A3.243 3.243 0 0 0 13.25 5.5Z"
        },
    ]

    function updateActive(e) {
        index = e.currentTarget.id.valueOf()
    }

    function select() {
        SelectFile().then(r => {
            console.log(r);
            games = r
        })
    }

    function repair(id, e) {
        switch (id) {
            case 0: {
                //id
                if (confirm("Problem: Game has no ID on it's name.\n- Older OPL will not boot it.\n- HDL won't be able to inject it if ZSO format.\n\nFind ID & rename? Please, confirm")) {
                    RepairFile(e.currentTarget.value, id).then(r => {
                        alert(r)
                    })
                } else {
                    alert("Cancelled")
                }
                break;
            }
            case 1: {
                //opl
                //FIXME Info needed...
                if (confirm("Problem: OPL supports only .iso and .zso extensions. Otherwise, they will not show up.\n\nRename extension? Please, confirm")) {
                    RepairFile(e.currentTarget.value, id).then(r => {
                        alert(r)
                    })
                } else {
                    alert("Cancelled")
                }
                break;
            }
            case 2: {
                //hdl
                //FIXME 2352...?
                if (confirm("Problem: File block size isn't 2048.\n\nZero-fill the file until its size is divisible by 2048? Please, confirm")) {
                    RepairFile(e.currentTarget.value, id).then(r => {
                        alert(r + ' bytes added to the file.')
                    })
                } else {
                    alert("Cancelled")
                }
                break;
            }
            case 3: {
                //hdl, ZSO without ID
                if (confirm("Problem: HDL unable to inject ZSO without ID.\n\nFind ID & rename? Please, confirm")) {
                    RepairFile(e.currentTarget.value, id).then(r => {
                        alert(r)
                    })
                } else {
                    alert("Cancelled")
                }
                break;
            }
        }
    }
</script>

<main>
    <div class="border-b border-gray-200 dark:border-gray-700 block">
        <ul class="inline-flex -mb-px text-sm font-medium text-center text-gray-500 dark:text-gray-400">
            {#each items as i}
                <li class="me-2">
                    <button
                            id={i.id.toString()}
                            on:click={updateActive}
                            class="{index == i.id ? 'active border-b-2 text-blue-500 border-blue-500' : 'hover:border-gray-300 hover:text-gray-300'} inline-flex items-center justify-center p-4 border-b-2 border-transparent rounded-t-lg group">
                        <svg class="w-4 h-4 me-2 text-gray-400 group-hover:text-gray-500 dark:text-gray-500 dark:group-hover:text-gray-300"
                             aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
                             viewBox="0 0 20 20">
                            <path d={i.v}/>
                        </svg>
                        {i.name}
                    </button>
                </li>
            {/each}
        </ul>
    </div>
    {#if index == 0}
        PS2 Page
    {:else if index == 1}
        <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-10" on:click={select}>
            Load
            Games
        </button>
        <p class="mt-1 text-sm text-gray-300" id="file_input_help">ISO or ZSO (Compressed)</p>
        <table class="w-full table-auto text-gray-400 bg-slate-800 mt-10">
            <thead class="bg-gray-700 text-gray-400">
            <tr>
                <th scope="col">Format</th>
                <th scope="col">ID</th>
                <th scope="col">Name</th>
                <th scope="col">Size</th>
                <th colspan="2">Compatibility</th>
                <th colspan="3">ZSO Info</th>
            </tr>
            <tr>
                <th colspan="4"/>
                <th class="w-1/12">OPL</th>
                <th class="w-1/12">HDL</th>
                <th colspan="3"/>
            </tr>
            <tr>
                <th colspan="6"/>
                <th>O. Size</th>
                <th>B. Size</th>
                <th>I. Shift</th>
            </tr>
            </thead>
            <tbody class="font-medium whitespace-nowrap text-white">
            {#each games as {format, url, id, name, size, opl, hdl, zso}}
                <tr class="border-b bg-gray-800 border-gray-700">
                    <td>
                        {#if format === 0}
                            ISO/DVD
                        {:else if format === 1}
                            ISO/CD
                        {:else if format === 2}
                            ZSO
                        {:else if format === 3}
                            ZSO/DVD
                        {:else if format === 4}
                            ZSO/CD
                        {:else}
                            <button value={url} class="hover:bg-slate-700 py-2 px-4 rounded"
                                    on:click={() => alert('Isn\'t a game!')}>
                                &#10060;
                            </button>
                        {/if}
                    </td>
                    <td>
                        {#if id.length !== 0}
                            {id}
                        {:else}
                            <button value={url} class="hover:bg-slate-700 py-2 px-4 rounded"
                                    on:click={(e)=>repair(0,e)}>
                                &#10060;
                            </button>
                        {/if}
                    </td>
                    <td>{name}</td>
                    <td>{size}</td>
                    <td>
                        {#if opl}&#9989;
                        {:else}
                            <button value={url} class="hover:bg-slate-700 py-2 px-4 rounded"
                                    on:click={(e)=>repair(1,e)}>
                                &#10060;
                            </button>
                        {/if}
                    </td>
                    <td>
                        {#if hdl}
                            {#if format === 2 || format === 3 || format === 4}
                                {#if id.length === 0}
                                    <button value={url} class="hover:bg-slate-700 py-2 px-4 rounded"
                                            on:click={(e)=>repair(3,e)}>
                                        &#10060;
                                    </button>
                                {:else}
                                    &#9989;
                                {/if}
                            {:else}
                                &#9989;
                            {/if}
                        {:else}
                            <button value={url} class="hover:bg-slate-700 py-2 px-4 rounded"
                                    on:click={(e)=>repair(2,e)}>
                                &#10060;
                            </button>
                        {/if}
                    </td>
                    <td>
                        {#if zso.is_zso}
                            {zso.orig_size}
                        {/if}
                    </td>
                    <td>
                        {#if zso.is_zso}
                            {zso.block_size}
                        {/if}
                    </td>
                    <td>
                        {#if zso.is_zso}
                            {zso.index_shift}
                        {/if}
                    </td>
                </tr>
            {/each}
            </tbody>
        </table>
    {:else if index == 2}
        POPS Page
    {:else if index == 3}
        OPL Page
    {:else if index == 4}
        SaveData Page
    {:else if index == 5}
        Device Page
    {/if}
</main>

<style>
    :root {
        font-family: 'Nunito', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
        Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
        background-color: rgba(27, 38, 54, 1);
        text-align: center;
        color: white;
    }

    main {
        text-align: center;
        padding: 1em;
        margin: 0 auto;
        height: 100vh;
    }
</style>
