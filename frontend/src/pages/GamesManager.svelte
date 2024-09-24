<script>
    import {games, select} from "../services/gameService.js";
    import FixButton from "../components/FixButton.svelte";
    import HDLCmdButton from "../components/HDLCmdButton.svelte";

    let gamesList;
    $: gamesList = $games;
    //TODO SoC pending!!! Diagnostic func needed to prevent diag via {#if}
</script>

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
    {#if gamesList != null}
        {#each gamesList as {format, path, id, name, size, opl, hdl, zso}}
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
                        <FixButton path={path} id={-1}/>
                    {/if}
                </td>
                <td>
                    {#if format !== -1}
                        {#if id.length !== 0}
                            {id}
                        {:else}
                            <FixButton path={path} id={0}/>
                        {/if}
                    {/if}
                </td>
                <td>{name}</td>
                <td>{size}</td>
                <td>
                    {#if format !== -1}
                        {#if opl}&#9989;
                        {:else}
                            <FixButton path={path} id={1}/>
                        {/if}
                    {/if}
                </td>
                <td>
                    {#if format !== -1}
                        {#if hdl}
                            {#if format === 2 || format === 3 || format === 4}
                                {#if id.length === 0}
                                    <FixButton path={path} id={3}/>
                                {:else}
                                    <HDLCmdButton path={path}/>
                                {/if}
                            {:else}
                                <HDLCmdButton path={path}/>
                            {/if}
                        {:else}
                            <FixButton path={path} id={2}/>
                        {/if}
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
    {/if}
    </tbody>
</table>