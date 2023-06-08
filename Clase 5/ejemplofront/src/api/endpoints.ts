const uri = "http://localhost:3000";

export function getCpu() {
    return fetch(`${uri}/getCpu`, {
        headers: { 'Content-Type': 'application/json' },
        method: 'GET',
    });
}