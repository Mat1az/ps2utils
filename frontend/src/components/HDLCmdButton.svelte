<script>
    import {getHDL} from "../services/gameService.js";
    export let path;

</script>

<button value={path} class="hover:bg-slate-700 py-2 px-4 rounded"
        on:click={()=>getHDL(path)}>
    &#9989;
</button>