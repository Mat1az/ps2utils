<script>
    import {fix} from "../services/gameService.js";
    export let path;
    export let id;

</script>

<button value={path} class="hover:bg-slate-700 py-2 px-4 rounded"
        on:click={()=>fix(path,id)}>
    &#10060;
</button>