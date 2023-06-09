// const uri = "http://localhost:3000";

export function getCpu() {
    return fetch(`/getCpu`, {
        headers: { 'Content-Type': 'application/json' },
        method: 'GET',
    });
}