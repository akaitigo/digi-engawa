import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { FuriganaText } from "./FuriganaText";

describe("FuriganaText", () => {
	it("renders plain text when no furigana provided", () => {
		render(<FuriganaText text="こんにちは" furigana="" />);
		expect(screen.getByText("こんにちは")).toBeDefined();
	});

	it("renders ruby text for furigana format", () => {
		const { container } = render(<FuriganaText text="電源を入れる" furigana="電源(でんげん)を入(い)れる" />);
		const rubyElements = container.querySelectorAll("ruby");
		expect(rubyElements.length).toBe(2);

		const rtElements = container.querySelectorAll("rt");
		expect(rtElements[0]?.textContent).toBe("でんげん");
		expect(rtElements[1]?.textContent).toBe("い");
	});

	it("handles text without any furigana markers", () => {
		render(<FuriganaText text="テスト" furigana="テスト" />);
		expect(screen.getByText("テスト")).toBeDefined();
	});
});
