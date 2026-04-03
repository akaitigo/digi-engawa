export const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";
export const WS_BASE = process.env["NEXT_PUBLIC_WS_URL"] ?? "ws://localhost:8080/ws";

export async function fetchAPI<T>(path: string, options?: RequestInit): Promise<T> {
	const res = await fetch(`${API_BASE}${path}`, {
		headers: { "Content-Type": "application/json" },
		...options,
	});

	if (!res.ok) {
		const body = await res.text().catch(() => "");
		throw new Error(`HTTP ${String(res.status)}: ${body || "request failed"}`);
	}

	return res.json() as Promise<T>;
}
