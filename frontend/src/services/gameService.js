import {GetGame, GetHDL, FixFile, SelectFile} from "../../wailsjs/go/services/GameService.js";
import {ClipboardSetText} from "../../wailsjs/runtime/runtime.js";
import {writable} from 'svelte/store';

//TODO SoC pending!!! Diagnostic func needed to prevent diag via {#if}

export const games = writable([]);

export function select() {
    SelectFile().then(r => {
        games.set(r);
    })
}

export function reloadGame(path) {
    GetGame(path).then(updatedGame => {
        games.update(current => {
            return current.map(game =>
                game.path === updatedGame.path ? { ...game, ...updatedGame } : game
            );
        })
    })
}

export function getHDL(path) {
    GetHDL(path).then(r => {
        if (confirm("Copy cmd to the clipboard? Change '/dev/sdx' to PS2 HDD location.\n\n" + r)) {
            ClipboardSetText(r).then(() => {
                alert("Copied to the clipboard.")
            }).catch(err => alert("Error: " + err))
        } else {
            alert("Cancelled")
        }
    })
    reloadGame(path);
}

export function fix(path,id) {
    //TODO Verify again after doing file changes
    switch (id) {
        case -1: {
            //isn't a game
            FixFile(path, id).then(r => {
                alert(r)
            })
            break;
        }
        case 0: {
            //id
            if (confirm("Problem: Game has no ID on it's name.\n- Older OPL will not boot it.\n- HDL won't be able to inject it if ZSO format.\n\nFind ID & rename? Please, confirm")) {
                FixFile(path, id).then(r => {
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
                FixFile(path, id).then(r => {
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
                FixFile(path, id).then(r => {
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
                FixFile(path, id).then(r => {
                    alert(r)
                })
            } else {
                alert("Cancelled")
            }
            break;
        }
    }
}