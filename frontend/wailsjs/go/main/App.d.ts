import {main} from '../models';

export function SelectFile():Promise<Array<main.Game>>;
export function RepairFile(path:string,id:int):Promise<string>;
export function GetHDL(path:string):Promise<string>;
export function GetGame(path:string):Promise<main.Game>;