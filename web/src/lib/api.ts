const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";

export async function fetchAPI<T>(path: string, options?: RequestInit): Promise<T> {
	const res = await fetch(`${API_BASE}${path}`, {
		headers: { "Content-Type": "application/json" },
		...options,
	});

	if (!res.ok) {
		const error = await res.json().catch(() => ({ error: "Unknown error" }));
		throw new Error((error as { error: string }).error ?? `HTTP ${String(res.status)}`);
	}

	return res.json() as Promise<T>;
}
