"use client";

import { useEffect, useState } from "react";
import { ClassroomCard } from "@/components/ClassroomCard";
import type { Classroom } from "@/types/classroom";

const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";

export default function DashboardPage() {
	const [classrooms, setClassrooms] = useState<Classroom[]>([]);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		const load = async () => {
			try {
				const res = await fetch(`${API_BASE}/api/classrooms`);
				if (res.ok) {
					const data = (await res.json()) as Classroom[];
					setClassrooms(data);
				}
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
