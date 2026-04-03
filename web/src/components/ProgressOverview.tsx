"use client";

import type { LearnerProgress } from "@/types/progress";

interface ProgressOverviewProps {
	progressList: LearnerProgress[];
	participantNames: Record<string, string>;
	totalSteps: number;
}

export function ProgressOverview({ progressList, participantNames, totalSteps }: ProgressOverviewProps) {
	if (progressList.length === 0) {
		return <p style={{ textAlign: "center", color: "#999", padding: "1rem" }}>進捗データはありません</p>;
	}

	return (
		<div>
			{progressList.map((p) => {
				const name = participantNames[p.participant_id] ?? "不明";
				const percentage = totalSteps > 0 ? Math.round((p.current_step / totalSteps) * 100) : 0;

				return (
					<div
						key={p.id}
						style={{
							padding: "0.75rem 1rem",
							borderRadius: "8px",
							backgroundColor: p.completed ? "#e8f5e9" : "#f5f5f5",
							marginBottom: "0.5rem",
							border: p.completed ? "2px solid #4caf50" : "1px solid #e0e0e0",
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
							<span style={{ fontSize: "1.125rem", fontWeight: "bold" }}>{name}</span>
							<span style={{ fontSize: "0.875rem", color: p.completed ? "#2e7d32" : "#666" }}>
								{p.completed ? "完了" : `${String(p.current_step)}/${String(totalSteps)} ステップ`}
							</span>
						</div>

						<div
							role="progressbar"
							aria-valuenow={percentage}
							aria-valuemin={0}
							aria-valuemax={100}
							aria-label={`${name}の進捗: ${String(percentage)}%`}
							style={{
								width: "100%",
								height: "8px",
								backgroundColor: "#e0e0e0",
								borderRadius: "4px",
								overflow: "hidden",
							}}
						>
							<div
								style={{
									width: `${String(percentage)}%`,
									height: "100%",
									backgroundColor: p.completed ? "#4caf50" : "#4A7C59",
									borderRadius: "4px",
									transition: "width 0.3s ease",
								}}
							/>
						</div>
					</div>
				);
			})}
		</div>
	);
}
