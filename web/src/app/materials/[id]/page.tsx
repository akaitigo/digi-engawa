"use client";

import { useCallback, useEffect, useState } from "react";
import { StepViewer } from "@/components/StepViewer";
import { fetchAPI } from "@/lib/api";
import type { Material } from "@/types/material";

export default function MaterialViewerPage({ params }: { params: Promise<{ id: string }> }) {
	const [material, setMaterial] = useState<Material | null>(null);
	const [error, setError] = useState<string | null>(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		let cancelled = false;
		const load = async () => {
			const { id } = await params;
			try {
				const data = await fetchAPI<Material>(`/api/materials/${id}`);
				if (!cancelled) setMaterial(data);
			} catch (e) {
				if (!cancelled) setError(e instanceof Error ? e.message : "読み込みエラー");
			} finally {
				if (!cancelled) setLoading(false);
			}
		};
		void load();
		return () => {
			cancelled = true;
		};
	}, [params]);

	const handleStepChange = useCallback((_stepOrder: number) => {
		// Progress tracking is handled via PUT /api/progress
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
