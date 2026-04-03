"use client";

import { useCallback } from "react";
import type { HelpRequest } from "@/types/help";

interface HelpRequestCardProps {
	request: HelpRequest;
	onStatusChange?: (id: string, status: string) => void;
}

const statusLabels: Record<string, string> = {
	pending: "🔴 たすけて",
	in_progress: "🟡 たいおうちゅう",
	resolved: "🟢 かいけつ",
};

const statusColors: Record<string, string> = {
	pending: "#ffebee",
	in_progress: "#fff3e0",
	resolved: "#e8f5e9",
};

export function HelpRequestCard({ request, onStatusChange }: HelpRequestCardProps) {
	const handleAction = useCallback(() => {
		if (!onStatusChange) return;
		if (request.status === "pending") {
			onStatusChange(request.id, "in_progress");
		} else if (request.status === "in_progress") {
			onStatusChange(request.id, "resolved");
		}
	}, [request, onStatusChange]);

	const nextAction =
		request.status === "pending" ? "たいおうする" : request.status === "in_progress" ? "かいけつ" : null;

	return (
		<div
			style={{
				padding: "1rem",
				borderRadius: "12px",
				backgroundColor: statusColors[request.status] ?? "#f5f5f5",
				border: "2px solid #e0e0e0",
				marginBottom: "0.75rem",
			}}
		>
			<div
				style={{
					display: "flex",
					justifyContent: "space-between",
					alignItems: "center",
					marginBottom: "0.5rem",
				}}
			>
				<span style={{ fontSize: "1.25rem", fontWeight: "bold" }}>
					{statusLabels[request.status] ?? request.status}
				</span>
				<span style={{ fontSize: "0.875rem", color: "#666" }}>
					{new Date(request.created_at).toLocaleTimeString("ja-JP")}
				</span>
			</div>

			{nextAction && (
				<button
					type="button"
					onClick={handleAction}
					style={{
						width: "100%",
						padding: "12px",
						fontSize: "1.125rem",
						fontWeight: "bold",
						minHeight: "48px",
						borderRadius: "8px",
						border: "2px solid #4A7C59",
						backgroundColor: "#4A7C59",
						color: "#ffffff",
						cursor: "pointer",
					}}
				>
					{nextAction}
				</button>
			)}
		</div>
	);
}
