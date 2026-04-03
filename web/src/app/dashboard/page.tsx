"use client";

import { useEffect, useState } from "react";
import { ClassroomCard } from "@/components/ClassroomCard";
import { fetchAPI } from "@/lib/api";
import type { Classroom } from "@/types/classroom";

export default function DashboardPage() {
	const [classrooms, setClassrooms] = useState<Classroom[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		const load = async () => {
			try {
				const data = await fetchAPI<Classroom[]>("/api/classrooms");
				setClassrooms(data);
			} catch (e) {
				setError(e instanceof Error ? e.message : "読み込みエラー");
			} finally {
				setLoading(false);
			}
		};
		void load();
	}, []);

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

	return (
		<main style={{ maxWidth: "600px", margin: "0 auto", padding: "1rem" }}>
			<div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: "1.5rem" }}>
				<h1 style={{ fontSize: "2rem", fontWeight: "bold" }}>教室一覧</h1>
				<a
					href="/dashboard/new"
					style={{
						padding: "12px 24px",
						fontSize: "1rem",
						fontWeight: "bold",
						borderRadius: "8px",
						backgroundColor: "#4A7C59",
						color: "#ffffff",
						textDecoration: "none",
					}}
				>
					+ 新しい教室
				</a>
			</div>

			{classrooms.length === 0 ? (
				<p style={{ textAlign: "center", fontSize: "1.125rem", color: "#999", padding: "2rem" }}>
					まだ教室がありません
				</p>
			) : (
				classrooms.map((c) => (
					<a key={c.id} href={`/dashboard/${c.id}`} style={{ textDecoration: "none" }}>
						<ClassroomCard classroom={c} />
					</a>
				))
			)}
		</main>
	);
}
