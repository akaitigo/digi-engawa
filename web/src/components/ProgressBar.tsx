"use client";

interface ProgressBarProps {
	current: number;
	total: number;
}

export function ProgressBar({ current, total }: ProgressBarProps) {
	const percentage = total > 0 ? Math.round((current / total) * 100) : 0;

	return (
		<div style={{ width: "100%", marginBottom: "1rem" }}>
			<div
				style={{
					display: "flex",
					justifyContent: "space-between",
					marginBottom: "0.5rem",
					fontSize: "1.125rem",
					fontWeight: "bold",
				}}
			>
				<span>
					{String(current)} / {String(total)} ステップ
				</span>
				<span>{String(percentage)}%</span>
			</div>
			<div
				role="progressbar"
				aria-valuenow={percentage}
				aria-valuemin={0}
				aria-valuemax={100}
				aria-label={`進捗: ${String(percentage)}%`}
				style={{
					width: "100%",
					height: "12px",
					backgroundColor: "#e0e0e0",
					borderRadius: "6px",
					overflow: "hidden",
				}}
			>
				<div
					style={{
						width: `${String(percentage)}%`,
						height: "100%",
						backgroundColor: "#4A7C59",
						borderRadius: "6px",
						transition: "width 0.3s ease",
					}}
				/>
			</div>
		</div>
	);
}
