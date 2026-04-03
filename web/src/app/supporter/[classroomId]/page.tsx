"use client";

import { use } from "react";
import { HelpBoard } from "@/components/HelpBoard";

export default function SupporterPage({ params }: { params: Promise<{ classroomId: string }> }) {
	const { classroomId } = use(params);

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
				サポーター画面
			</h1>
			<HelpBoard classroomId={classroomId} />
		</main>
	);
}
