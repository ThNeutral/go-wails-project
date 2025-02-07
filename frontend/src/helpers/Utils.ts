export function debounce(callback: () => void, options: {
    timeout: number,
    before?: () => void
}) {
    let timeoutID: number;
    let hasCalledBefore = false;
    return function () {
        if (!hasCalledBefore && options.before) {
            options.before()
        }
        if (timeoutID !== undefined) clearTimeout(timeoutID)

        timeoutID = setTimeout(() => {
            hasCalledBefore = false; 
            callback();
        }, options.timeout);
    }       
}

export function isErrorObject(obj: any): {Error: string} | null {
    if (obj && obj.Error) {
        return obj
    }
    return null
}

export function captialize(str: string) {
    return str.split("").map((char, index) => (index === 0 ? char.toUpperCase() : char.toLowerCase())).join("")
}