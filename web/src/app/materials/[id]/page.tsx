"use client";

import { useCallback, useEffect, useState } from "react";
import { StepViewer } from "@/components/StepViewer";
import type { Material } from "@/types/material";

const API_BASE = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

export default function MaterialViewerPage({ params }: { params: Promise<{ id: string }> }) {
	const [material, setMaterial] = useState<Material | null>(null);
	const [error, setError] = useState<string | null>(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		const load = async () => {
			const { id } = await params;
			try {
				const res = await fetch(`${API_BASE}/api/materials/${id}`);
				if (!res.ok) {
					throw new Error("教材が見つかりません");
				}
				const data = (await res.json()) as Material;
				setMaterial(data);
			} catch (e) {
				setError(e instanceof Error ? e.message : "読み込みエラー");
			} finally {
				setLoading(false);
			}
		};
		void load();
	}, [params]);

	const handleStepChange = useCallback(
		(stepOrder: number) => {
			if (!material) return;
			// Progress tracking will be added in Issue #8
			void stepOrder;
		},
		[material],
	);

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

	if (!material) {
		return null;
	}

	return (
		<main style={{ padding: "1rem" }}>
			<h1
				style={{
					fontSize: "2rem",
					fontWeight: "bold",
					textAlign: "center",
					marginBottom: "1rem",
					color: "#333",
				}}
			>
				{material.title}
			</h1>
			<StepViewer steps={material.steps} onStepChange={handleStepChange} />
		</main>
	);
}
