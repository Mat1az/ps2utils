export function SelectFile() {
    return window['go']['services']['GameService']['SelectFile']();
}

export function FixFile(path, id) {
    return window['go']['services']['GameService']['FixFile'](path, id);
}

export function GetHDL(path) {
    return window['go']['services']['GameService']['GetHDL'](path);
}

export function GetGame(path) {
    return window['go']['services']['GameService']['GetGame'](path);
}