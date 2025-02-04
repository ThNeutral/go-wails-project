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