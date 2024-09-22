export function SelectFile() {
    return window['go']['main']['App']['SelectFile']();
}

export function RepairFile(path, id) {
    return window['go']['main']['App']['RepairFile'](path, id);
}

export function GetHDL(path) {
    return window['go']['main']['App']['GetHDL'](path);
}

export function GetGame(path) {
    return window['go']['main']['App']['GetGame'](path);
}