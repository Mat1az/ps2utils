import {models} from '../models';

export function SelectFile():Promise<Array<models.Game>>;
export function FixFile(path:string,id:int):Promise<string>;
export function GetHDL(path:string):Promise<string>;
export function GetGame(path:string):Promise<models.Game>;