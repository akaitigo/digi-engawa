"use client";

import { use, useEffect, useState } from "react";
import type { Classroom } from "@/types/classroom";

const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";

export default function JoinPage({ params }: { params: Promise<{ code: string }> }) {
	const { code } = use(params);
	const [classroom, setClassroom] = useState<Classroom | null>(null);
	const [error, setError] = useState<string | null>(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		const load = async () => {
			try {
				const res = await fetch(`${API_BASE}/api/join/${code}`);
				if (res.ok) {
					setClassroom((await res.json()) as Classroom);
				} else {
					setError("教室が見つかりません");
				}
			} catch {
				setError("読み込みエラー");
			} finally {
				setLoading(false);
			}
		};
		void load();
	}, [code]);

	if (loading) {
		return (
			<main style={{ padding: "2rem", textAlign: "center" }}>
				<p style={{ fontSize: "1.5rem" }}>よみこみちゅう...</p>
			</main>
		);
	}

	if (error) {
		return (
			<main style={{ padding: "2rem", textAlign: "center" }}>
				<p style={{ fontSize: "1.5rem", color: "#d32f2f" }}>{error}</p>
			</main>
		);
	}

	if (!classroom) return null;

	return (
		<main style={{ maxWidth: "500px", margin: "0 auto", padding: "2rem", textAlign: "center" }}>
			<h1 style={{ fontSize: "2rem", fontWeight: "bold", marginBottom: "1rem" }}>{classroom.title}</h1>
			<p style={{ fontSize: "1.125rem", color: "#666", marginBottom: "1.5rem" }}>{classroom.description}</p>
			<p style={{ fontSize: "1rem", color: "#666" }}>
				{new Date(classroom.scheduled_at).toLocaleString("ja-JP")} @ {classroom.location}
			</p>
		</main>
	);
}
