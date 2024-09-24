import { writable } from 'svelte/store';
import {GetGame} from "../../wailsjs/go/main/App.js";

const games = writable([]);

games.subscribe((value) => {
    console.log(value);
}); // logs '0'

games.set(1); // logs '1'

games.update((n) => n + 1); // logs '2'