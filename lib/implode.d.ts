export declare type ChainFile = {
    chain: Link[];
};
declare type Link = {
    seqno: number;
    kid: string | null;
    sig: string;
    payload_json: string | null;
    sig_version: number | null;
    merkle_seqno: number | null;
    si1: string | null;
    si2: string | null;
    s2: string | null;
};
export declare const implode: (f: ChainFile) => ChainFile;
export declare const main: () => Promise<void>;
export {};
