import {main} from '../models';

export function SelectFile():Promise<Array<main.Game>>;
export function RepairFile(path:string,id:int):Promise<string>;