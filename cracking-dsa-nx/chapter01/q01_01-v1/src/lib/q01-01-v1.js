/**
 * JavaScript implementation (v1) for Chapter 01, Interview Question 01 of Cracking the Coding Interview
 */
export function isUnique_ForOfWithSet(maybeUnique) {
    const set = new Set();
    for(const c of maybeUnique) {
        if(set.has(c)) return false;
        set.add(c);
    }
    return true;
}

export function isUnique_OldForWithSet(maybeUnique) {
    const set = new Set();
    for(let i = 0; i < maybeUnique.length; i++) {
        if(set.has(maybeUnique[i])) return false
        set.add(maybeUnique[i]);
    }
    return true;
};

export function isUnique_OldForNoDataStructures(maybeUnique) {
    for(let i = 0; i < maybeUnique.length - 1; i++) {
        for(let j = i+1; j < maybeUnique.length; j++) {
            console.log(`${maybeUnique[i]} : ${maybeUnique[j]}`)
            if(maybeUnique[i] === maybeUnique[j]) return false
        }
    }
    return true;
};

