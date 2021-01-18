/*
 * This file is part of "SnipSnap Radeox Rendering Engine".
 *
 * Copyright (c) 2002 Stephan J. Schmidt, Matthias L. Jugel
 * All Rights Reserved.
 *
 * Please visit http://radeox.org/ for updates and contact.
 *
 * --LICENSE NOTICE--
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * --LICENSE NOTICE--
 */

package org.radeox.filter.regex;

import java.util.Locale;
import java.util.ResourceBundle;

import lombok.extern.slf4j.Slf4j;
import org.radeox.api.engine.context.InitialRenderContext;
import org.radeox.api.engine.context.RenderContext;

/*
 * Class that extends RegexReplaceFilter but reads patterns from a locale file
 * instead of hardwired regex @author stephan @team sonicteam
 * 
 * @version $Id: LocaleRegexReplaceFilter.java 7707 2006-04-12 17:30:19Z
 *          ian@caret.cam.ac.uk $
 */
@Slf4j
public abstract class LocaleRegexReplaceFilter extends RegexReplaceFilter
{

	protected abstract String getLocaleKey();

	protected boolean isSingleLine()
	{
		return false;
	}

	protected ResourceBundle getInputBundle()
	{
		Locale inputLocale = (Locale) initialContext
				.get(RenderContext.INPUT_LOCALE);
		String inputName = (String) initialContext
				.get(RenderContext.INPUT_BUNDLE_NAME);
		return ResourceBundle.getBundle(inputName, inputLocale);
	}

	protected ResourceBundle getOutputBundle()
	{
		String outputName = (String) initialContext
				.get(RenderContext.OUTPUT_BUNDLE_NAME);
		Locale outputLocale = (Locale) initialContext
				.get(RenderContext.OUTPUT_LOCALE);
		return ResourceBundle.getBundle(outputName, outputLocale);
	}

	public void setInitialContext(InitialRenderContext context)
	{
		super.setInitialContext(context);
		clearRegex();

		ResourceBundle outputMessages = getOutputBundle();
		ResourceBundle inputMessages = getInputBundle();

		String match = inputMessages.getString(getLocaleKey() + ".match");
		String print = outputMessages.getString(getLocaleKey() + ".print");
		// log.info(getLocaleKey()+": match="+match+"
		// pattern="+print);
		addRegex(match, print, isSingleLine() ? RegexReplaceFilter.SINGLELINE
				: RegexReplaceFilter.MULTILINE);
	}
}