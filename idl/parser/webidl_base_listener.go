// Code generated from WebIDL by ANTLR 4.7.2. DO NOT EDIT.

package parser // WebIDL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWebIDLListener is a complete listener for a parse tree produced by WebIDLParser.
type BaseWebIDLListener struct{}

var _ WebIDLListener = &BaseWebIDLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWebIDLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWebIDLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWebIDLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWebIDLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterWebIDL is called when production webIDL is entered.
func (s *BaseWebIDLListener) EnterWebIDL(ctx *WebIDLContext) {}

// ExitWebIDL is called when production webIDL is exited.
func (s *BaseWebIDLListener) ExitWebIDL(ctx *WebIDLContext) {}

// EnterDefinitions is called when production definitions is entered.
func (s *BaseWebIDLListener) EnterDefinitions(ctx *DefinitionsContext) {}

// ExitDefinitions is called when production definitions is exited.
func (s *BaseWebIDLListener) ExitDefinitions(ctx *DefinitionsContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseWebIDLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseWebIDLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterCallbackOrInterface is called when production callbackOrInterface is entered.
func (s *BaseWebIDLListener) EnterCallbackOrInterface(ctx *CallbackOrInterfaceContext) {}

// ExitCallbackOrInterface is called when production callbackOrInterface is exited.
func (s *BaseWebIDLListener) ExitCallbackOrInterface(ctx *CallbackOrInterfaceContext) {}

// EnterCallbackRestOrInterface is called when production callbackRestOrInterface is entered.
func (s *BaseWebIDLListener) EnterCallbackRestOrInterface(ctx *CallbackRestOrInterfaceContext) {}

// ExitCallbackRestOrInterface is called when production callbackRestOrInterface is exited.
func (s *BaseWebIDLListener) ExitCallbackRestOrInterface(ctx *CallbackRestOrInterfaceContext) {}

// EnterInterface_ is called when production interface_ is entered.
func (s *BaseWebIDLListener) EnterInterface_(ctx *Interface_Context) {}

// ExitInterface_ is called when production interface_ is exited.
func (s *BaseWebIDLListener) ExitInterface_(ctx *Interface_Context) {}

// EnterClass_ is called when production class_ is entered.
func (s *BaseWebIDLListener) EnterClass_(ctx *Class_Context) {}

// ExitClass_ is called when production class_ is exited.
func (s *BaseWebIDLListener) ExitClass_(ctx *Class_Context) {}

// EnterPartial is called when production partial is entered.
func (s *BaseWebIDLListener) EnterPartial(ctx *PartialContext) {}

// ExitPartial is called when production partial is exited.
func (s *BaseWebIDLListener) ExitPartial(ctx *PartialContext) {}

// EnterPartialDefinition is called when production partialDefinition is entered.
func (s *BaseWebIDLListener) EnterPartialDefinition(ctx *PartialDefinitionContext) {}

// ExitPartialDefinition is called when production partialDefinition is exited.
func (s *BaseWebIDLListener) ExitPartialDefinition(ctx *PartialDefinitionContext) {}

// EnterPartialInterface is called when production partialInterface is entered.
func (s *BaseWebIDLListener) EnterPartialInterface(ctx *PartialInterfaceContext) {}

// ExitPartialInterface is called when production partialInterface is exited.
func (s *BaseWebIDLListener) ExitPartialInterface(ctx *PartialInterfaceContext) {}

// EnterInterfaceMembers is called when production interfaceMembers is entered.
func (s *BaseWebIDLListener) EnterInterfaceMembers(ctx *InterfaceMembersContext) {}

// ExitInterfaceMembers is called when production interfaceMembers is exited.
func (s *BaseWebIDLListener) ExitInterfaceMembers(ctx *InterfaceMembersContext) {}

// EnterInterfaceMember is called when production interfaceMember is entered.
func (s *BaseWebIDLListener) EnterInterfaceMember(ctx *InterfaceMemberContext) {}

// ExitInterfaceMember is called when production interfaceMember is exited.
func (s *BaseWebIDLListener) ExitInterfaceMember(ctx *InterfaceMemberContext) {}

// EnterDictionary is called when production dictionary is entered.
func (s *BaseWebIDLListener) EnterDictionary(ctx *DictionaryContext) {}

// ExitDictionary is called when production dictionary is exited.
func (s *BaseWebIDLListener) ExitDictionary(ctx *DictionaryContext) {}

// EnterDictionaryMembers is called when production dictionaryMembers is entered.
func (s *BaseWebIDLListener) EnterDictionaryMembers(ctx *DictionaryMembersContext) {}

// ExitDictionaryMembers is called when production dictionaryMembers is exited.
func (s *BaseWebIDLListener) ExitDictionaryMembers(ctx *DictionaryMembersContext) {}

// EnterDictionaryMember is called when production dictionaryMember is entered.
func (s *BaseWebIDLListener) EnterDictionaryMember(ctx *DictionaryMemberContext) {}

// ExitDictionaryMember is called when production dictionaryMember is exited.
func (s *BaseWebIDLListener) ExitDictionaryMember(ctx *DictionaryMemberContext) {}

// EnterRequired is called when production required is entered.
func (s *BaseWebIDLListener) EnterRequired(ctx *RequiredContext) {}

// ExitRequired is called when production required is exited.
func (s *BaseWebIDLListener) ExitRequired(ctx *RequiredContext) {}

// EnterPartialDictionary is called when production partialDictionary is entered.
func (s *BaseWebIDLListener) EnterPartialDictionary(ctx *PartialDictionaryContext) {}

// ExitPartialDictionary is called when production partialDictionary is exited.
func (s *BaseWebIDLListener) ExitPartialDictionary(ctx *PartialDictionaryContext) {}

// EnterDefault_ is called when production default_ is entered.
func (s *BaseWebIDLListener) EnterDefault_(ctx *Default_Context) {}

// ExitDefault_ is called when production default_ is exited.
func (s *BaseWebIDLListener) ExitDefault_(ctx *Default_Context) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseWebIDLListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseWebIDLListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterInheritance is called when production inheritance is entered.
func (s *BaseWebIDLListener) EnterInheritance(ctx *InheritanceContext) {}

// ExitInheritance is called when production inheritance is exited.
func (s *BaseWebIDLListener) ExitInheritance(ctx *InheritanceContext) {}

// EnterExtension is called when production extension is entered.
func (s *BaseWebIDLListener) EnterExtension(ctx *ExtensionContext) {}

// ExitExtension is called when production extension is exited.
func (s *BaseWebIDLListener) ExitExtension(ctx *ExtensionContext) {}

// EnterEnum_ is called when production enum_ is entered.
func (s *BaseWebIDLListener) EnterEnum_(ctx *Enum_Context) {}

// ExitEnum_ is called when production enum_ is exited.
func (s *BaseWebIDLListener) ExitEnum_(ctx *Enum_Context) {}

// EnterEnumValueList is called when production enumValueList is entered.
func (s *BaseWebIDLListener) EnterEnumValueList(ctx *EnumValueListContext) {}

// ExitEnumValueList is called when production enumValueList is exited.
func (s *BaseWebIDLListener) ExitEnumValueList(ctx *EnumValueListContext) {}

// EnterEnumValueListComma is called when production enumValueListComma is entered.
func (s *BaseWebIDLListener) EnterEnumValueListComma(ctx *EnumValueListCommaContext) {}

// ExitEnumValueListComma is called when production enumValueListComma is exited.
func (s *BaseWebIDLListener) ExitEnumValueListComma(ctx *EnumValueListCommaContext) {}

// EnterEnumValueListString is called when production enumValueListString is entered.
func (s *BaseWebIDLListener) EnterEnumValueListString(ctx *EnumValueListStringContext) {}

// ExitEnumValueListString is called when production enumValueListString is exited.
func (s *BaseWebIDLListener) ExitEnumValueListString(ctx *EnumValueListStringContext) {}

// EnterCallbackRest is called when production callbackRest is entered.
func (s *BaseWebIDLListener) EnterCallbackRest(ctx *CallbackRestContext) {}

// ExitCallbackRest is called when production callbackRest is exited.
func (s *BaseWebIDLListener) ExitCallbackRest(ctx *CallbackRestContext) {}

// EnterTypedef is called when production typedef is entered.
func (s *BaseWebIDLListener) EnterTypedef(ctx *TypedefContext) {}

// ExitTypedef is called when production typedef is exited.
func (s *BaseWebIDLListener) ExitTypedef(ctx *TypedefContext) {}

// EnterImplementsStatement is called when production implementsStatement is entered.
func (s *BaseWebIDLListener) EnterImplementsStatement(ctx *ImplementsStatementContext) {}

// ExitImplementsStatement is called when production implementsStatement is exited.
func (s *BaseWebIDLListener) ExitImplementsStatement(ctx *ImplementsStatementContext) {}

// EnterConst_ is called when production const_ is entered.
func (s *BaseWebIDLListener) EnterConst_(ctx *Const_Context) {}

// ExitConst_ is called when production const_ is exited.
func (s *BaseWebIDLListener) ExitConst_(ctx *Const_Context) {}

// EnterConstValue is called when production constValue is entered.
func (s *BaseWebIDLListener) EnterConstValue(ctx *ConstValueContext) {}

// ExitConstValue is called when production constValue is exited.
func (s *BaseWebIDLListener) ExitConstValue(ctx *ConstValueContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseWebIDLListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseWebIDLListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseWebIDLListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseWebIDLListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterSerializer is called when production serializer is entered.
func (s *BaseWebIDLListener) EnterSerializer(ctx *SerializerContext) {}

// ExitSerializer is called when production serializer is exited.
func (s *BaseWebIDLListener) ExitSerializer(ctx *SerializerContext) {}

// EnterSerializerRest is called when production serializerRest is entered.
func (s *BaseWebIDLListener) EnterSerializerRest(ctx *SerializerRestContext) {}

// ExitSerializerRest is called when production serializerRest is exited.
func (s *BaseWebIDLListener) ExitSerializerRest(ctx *SerializerRestContext) {}

// EnterSerializationPattern is called when production serializationPattern is entered.
func (s *BaseWebIDLListener) EnterSerializationPattern(ctx *SerializationPatternContext) {}

// ExitSerializationPattern is called when production serializationPattern is exited.
func (s *BaseWebIDLListener) ExitSerializationPattern(ctx *SerializationPatternContext) {}

// EnterSerializationPatternMap is called when production serializationPatternMap is entered.
func (s *BaseWebIDLListener) EnterSerializationPatternMap(ctx *SerializationPatternMapContext) {}

// ExitSerializationPatternMap is called when production serializationPatternMap is exited.
func (s *BaseWebIDLListener) ExitSerializationPatternMap(ctx *SerializationPatternMapContext) {}

// EnterSerializationPatternList is called when production serializationPatternList is entered.
func (s *BaseWebIDLListener) EnterSerializationPatternList(ctx *SerializationPatternListContext) {}

// ExitSerializationPatternList is called when production serializationPatternList is exited.
func (s *BaseWebIDLListener) ExitSerializationPatternList(ctx *SerializationPatternListContext) {}

// EnterStringifier is called when production stringifier is entered.
func (s *BaseWebIDLListener) EnterStringifier(ctx *StringifierContext) {}

// ExitStringifier is called when production stringifier is exited.
func (s *BaseWebIDLListener) ExitStringifier(ctx *StringifierContext) {}

// EnterStringifierRest is called when production stringifierRest is entered.
func (s *BaseWebIDLListener) EnterStringifierRest(ctx *StringifierRestContext) {}

// ExitStringifierRest is called when production stringifierRest is exited.
func (s *BaseWebIDLListener) ExitStringifierRest(ctx *StringifierRestContext) {}

// EnterStaticMember is called when production staticMember is entered.
func (s *BaseWebIDLListener) EnterStaticMember(ctx *StaticMemberContext) {}

// ExitStaticMember is called when production staticMember is exited.
func (s *BaseWebIDLListener) ExitStaticMember(ctx *StaticMemberContext) {}

// EnterStaticMemberRest is called when production staticMemberRest is entered.
func (s *BaseWebIDLListener) EnterStaticMemberRest(ctx *StaticMemberRestContext) {}

// ExitStaticMemberRest is called when production staticMemberRest is exited.
func (s *BaseWebIDLListener) ExitStaticMemberRest(ctx *StaticMemberRestContext) {}

// EnterReadonlyMember is called when production readonlyMember is entered.
func (s *BaseWebIDLListener) EnterReadonlyMember(ctx *ReadonlyMemberContext) {}

// ExitReadonlyMember is called when production readonlyMember is exited.
func (s *BaseWebIDLListener) ExitReadonlyMember(ctx *ReadonlyMemberContext) {}

// EnterReadonlyMemberRest is called when production readonlyMemberRest is entered.
func (s *BaseWebIDLListener) EnterReadonlyMemberRest(ctx *ReadonlyMemberRestContext) {}

// ExitReadonlyMemberRest is called when production readonlyMemberRest is exited.
func (s *BaseWebIDLListener) ExitReadonlyMemberRest(ctx *ReadonlyMemberRestContext) {}

// EnterReadWriteAttribute is called when production readWriteAttribute is entered.
func (s *BaseWebIDLListener) EnterReadWriteAttribute(ctx *ReadWriteAttributeContext) {}

// ExitReadWriteAttribute is called when production readWriteAttribute is exited.
func (s *BaseWebIDLListener) ExitReadWriteAttribute(ctx *ReadWriteAttributeContext) {}

// EnterAttributeRest is called when production attributeRest is entered.
func (s *BaseWebIDLListener) EnterAttributeRest(ctx *AttributeRestContext) {}

// ExitAttributeRest is called when production attributeRest is exited.
func (s *BaseWebIDLListener) ExitAttributeRest(ctx *AttributeRestContext) {}

// EnterAttributeName is called when production attributeName is entered.
func (s *BaseWebIDLListener) EnterAttributeName(ctx *AttributeNameContext) {}

// ExitAttributeName is called when production attributeName is exited.
func (s *BaseWebIDLListener) ExitAttributeName(ctx *AttributeNameContext) {}

// EnterAttributeNameKeyword is called when production attributeNameKeyword is entered.
func (s *BaseWebIDLListener) EnterAttributeNameKeyword(ctx *AttributeNameKeywordContext) {}

// ExitAttributeNameKeyword is called when production attributeNameKeyword is exited.
func (s *BaseWebIDLListener) ExitAttributeNameKeyword(ctx *AttributeNameKeywordContext) {}

// EnterInherit is called when production inherit is entered.
func (s *BaseWebIDLListener) EnterInherit(ctx *InheritContext) {}

// ExitInherit is called when production inherit is exited.
func (s *BaseWebIDLListener) ExitInherit(ctx *InheritContext) {}

// EnterReadOnly is called when production readOnly is entered.
func (s *BaseWebIDLListener) EnterReadOnly(ctx *ReadOnlyContext) {}

// ExitReadOnly is called when production readOnly is exited.
func (s *BaseWebIDLListener) ExitReadOnly(ctx *ReadOnlyContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseWebIDLListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseWebIDLListener) ExitOperation(ctx *OperationContext) {}

// EnterSpecialOperation is called when production specialOperation is entered.
func (s *BaseWebIDLListener) EnterSpecialOperation(ctx *SpecialOperationContext) {}

// ExitSpecialOperation is called when production specialOperation is exited.
func (s *BaseWebIDLListener) ExitSpecialOperation(ctx *SpecialOperationContext) {}

// EnterSpecials is called when production specials is entered.
func (s *BaseWebIDLListener) EnterSpecials(ctx *SpecialsContext) {}

// ExitSpecials is called when production specials is exited.
func (s *BaseWebIDLListener) ExitSpecials(ctx *SpecialsContext) {}

// EnterSpecial is called when production special is entered.
func (s *BaseWebIDLListener) EnterSpecial(ctx *SpecialContext) {}

// ExitSpecial is called when production special is exited.
func (s *BaseWebIDLListener) ExitSpecial(ctx *SpecialContext) {}

// EnterOperationRest is called when production operationRest is entered.
func (s *BaseWebIDLListener) EnterOperationRest(ctx *OperationRestContext) {}

// ExitOperationRest is called when production operationRest is exited.
func (s *BaseWebIDLListener) ExitOperationRest(ctx *OperationRestContext) {}

// EnterOptionalIdentifier is called when production optionalIdentifier is entered.
func (s *BaseWebIDLListener) EnterOptionalIdentifier(ctx *OptionalIdentifierContext) {}

// ExitOptionalIdentifier is called when production optionalIdentifier is exited.
func (s *BaseWebIDLListener) ExitOptionalIdentifier(ctx *OptionalIdentifierContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseWebIDLListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseWebIDLListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseWebIDLListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseWebIDLListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseWebIDLListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseWebIDLListener) ExitArgument(ctx *ArgumentContext) {}

// EnterOptionalOrRequiredArgument is called when production optionalOrRequiredArgument is entered.
func (s *BaseWebIDLListener) EnterOptionalOrRequiredArgument(ctx *OptionalOrRequiredArgumentContext) {}

// ExitOptionalOrRequiredArgument is called when production optionalOrRequiredArgument is exited.
func (s *BaseWebIDLListener) ExitOptionalOrRequiredArgument(ctx *OptionalOrRequiredArgumentContext) {}

// EnterArgumentName is called when production argumentName is entered.
func (s *BaseWebIDLListener) EnterArgumentName(ctx *ArgumentNameContext) {}

// ExitArgumentName is called when production argumentName is exited.
func (s *BaseWebIDLListener) ExitArgumentName(ctx *ArgumentNameContext) {}

// EnterEllipsis is called when production ellipsis is entered.
func (s *BaseWebIDLListener) EnterEllipsis(ctx *EllipsisContext) {}

// ExitEllipsis is called when production ellipsis is exited.
func (s *BaseWebIDLListener) ExitEllipsis(ctx *EllipsisContext) {}

// EnterIterable is called when production iterable is entered.
func (s *BaseWebIDLListener) EnterIterable(ctx *IterableContext) {}

// ExitIterable is called when production iterable is exited.
func (s *BaseWebIDLListener) ExitIterable(ctx *IterableContext) {}

// EnterOptionalType is called when production optionalType is entered.
func (s *BaseWebIDLListener) EnterOptionalType(ctx *OptionalTypeContext) {}

// ExitOptionalType is called when production optionalType is exited.
func (s *BaseWebIDLListener) ExitOptionalType(ctx *OptionalTypeContext) {}

// EnterReadWriteMaplike is called when production readWriteMaplike is entered.
func (s *BaseWebIDLListener) EnterReadWriteMaplike(ctx *ReadWriteMaplikeContext) {}

// ExitReadWriteMaplike is called when production readWriteMaplike is exited.
func (s *BaseWebIDLListener) ExitReadWriteMaplike(ctx *ReadWriteMaplikeContext) {}

// EnterReadWriteSetlike is called when production readWriteSetlike is entered.
func (s *BaseWebIDLListener) EnterReadWriteSetlike(ctx *ReadWriteSetlikeContext) {}

// ExitReadWriteSetlike is called when production readWriteSetlike is exited.
func (s *BaseWebIDLListener) ExitReadWriteSetlike(ctx *ReadWriteSetlikeContext) {}

// EnterMaplikeRest is called when production maplikeRest is entered.
func (s *BaseWebIDLListener) EnterMaplikeRest(ctx *MaplikeRestContext) {}

// ExitMaplikeRest is called when production maplikeRest is exited.
func (s *BaseWebIDLListener) ExitMaplikeRest(ctx *MaplikeRestContext) {}

// EnterSetlikeRest is called when production setlikeRest is entered.
func (s *BaseWebIDLListener) EnterSetlikeRest(ctx *SetlikeRestContext) {}

// ExitSetlikeRest is called when production setlikeRest is exited.
func (s *BaseWebIDLListener) ExitSetlikeRest(ctx *SetlikeRestContext) {}

// EnterExtendedAttributeList is called when production extendedAttributeList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeList(ctx *ExtendedAttributeListContext) {}

// ExitExtendedAttributeList is called when production extendedAttributeList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeList(ctx *ExtendedAttributeListContext) {}

// EnterExtendedAttributes is called when production extendedAttributes is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributes(ctx *ExtendedAttributesContext) {}

// ExitExtendedAttributes is called when production extendedAttributes is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributes(ctx *ExtendedAttributesContext) {}

// EnterExtendedAttribute is called when production extendedAttribute is entered.
func (s *BaseWebIDLListener) EnterExtendedAttribute(ctx *ExtendedAttributeContext) {}

// ExitExtendedAttribute is called when production extendedAttribute is exited.
func (s *BaseWebIDLListener) ExitExtendedAttribute(ctx *ExtendedAttributeContext) {}

// EnterExtendedAttributeRest is called when production extendedAttributeRest is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeRest(ctx *ExtendedAttributeRestContext) {}

// ExitExtendedAttributeRest is called when production extendedAttributeRest is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeRest(ctx *ExtendedAttributeRestContext) {}

// EnterExtendedAttributeInner is called when production extendedAttributeInner is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeInner(ctx *ExtendedAttributeInnerContext) {}

// ExitExtendedAttributeInner is called when production extendedAttributeInner is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeInner(ctx *ExtendedAttributeInnerContext) {}

// EnterOther is called when production other is entered.
func (s *BaseWebIDLListener) EnterOther(ctx *OtherContext) {}

// ExitOther is called when production other is exited.
func (s *BaseWebIDLListener) ExitOther(ctx *OtherContext) {}

// EnterArgumentNameKeyword is called when production argumentNameKeyword is entered.
func (s *BaseWebIDLListener) EnterArgumentNameKeyword(ctx *ArgumentNameKeywordContext) {}

// ExitArgumentNameKeyword is called when production argumentNameKeyword is exited.
func (s *BaseWebIDLListener) ExitArgumentNameKeyword(ctx *ArgumentNameKeywordContext) {}

// EnterOtherOrComma is called when production otherOrComma is entered.
func (s *BaseWebIDLListener) EnterOtherOrComma(ctx *OtherOrCommaContext) {}

// ExitOtherOrComma is called when production otherOrComma is exited.
func (s *BaseWebIDLListener) ExitOtherOrComma(ctx *OtherOrCommaContext) {}

// EnterType is called when production type is entered.
func (s *BaseWebIDLListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseWebIDLListener) ExitType(ctx *TypeContext) {}

// EnterSingleType is called when production singleType is entered.
func (s *BaseWebIDLListener) EnterSingleType(ctx *SingleTypeContext) {}

// ExitSingleType is called when production singleType is exited.
func (s *BaseWebIDLListener) ExitSingleType(ctx *SingleTypeContext) {}

// EnterUnionType is called when production unionType is entered.
func (s *BaseWebIDLListener) EnterUnionType(ctx *UnionTypeContext) {}

// ExitUnionType is called when production unionType is exited.
func (s *BaseWebIDLListener) ExitUnionType(ctx *UnionTypeContext) {}

// EnterUnionMemberType is called when production unionMemberType is entered.
func (s *BaseWebIDLListener) EnterUnionMemberType(ctx *UnionMemberTypeContext) {}

// ExitUnionMemberType is called when production unionMemberType is exited.
func (s *BaseWebIDLListener) ExitUnionMemberType(ctx *UnionMemberTypeContext) {}

// EnterUnionMemberTypes is called when production unionMemberTypes is entered.
func (s *BaseWebIDLListener) EnterUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// ExitUnionMemberTypes is called when production unionMemberTypes is exited.
func (s *BaseWebIDLListener) ExitUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// EnterNonAnyType is called when production nonAnyType is entered.
func (s *BaseWebIDLListener) EnterNonAnyType(ctx *NonAnyTypeContext) {}

// ExitNonAnyType is called when production nonAnyType is exited.
func (s *BaseWebIDLListener) ExitNonAnyType(ctx *NonAnyTypeContext) {}

// EnterBufferRelatedType is called when production bufferRelatedType is entered.
func (s *BaseWebIDLListener) EnterBufferRelatedType(ctx *BufferRelatedTypeContext) {}

// ExitBufferRelatedType is called when production bufferRelatedType is exited.
func (s *BaseWebIDLListener) ExitBufferRelatedType(ctx *BufferRelatedTypeContext) {}

// EnterConstType is called when production constType is entered.
func (s *BaseWebIDLListener) EnterConstType(ctx *ConstTypeContext) {}

// ExitConstType is called when production constType is exited.
func (s *BaseWebIDLListener) ExitConstType(ctx *ConstTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseWebIDLListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseWebIDLListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterUnrestrictedFloatType is called when production unrestrictedFloatType is entered.
func (s *BaseWebIDLListener) EnterUnrestrictedFloatType(ctx *UnrestrictedFloatTypeContext) {}

// ExitUnrestrictedFloatType is called when production unrestrictedFloatType is exited.
func (s *BaseWebIDLListener) ExitUnrestrictedFloatType(ctx *UnrestrictedFloatTypeContext) {}

// EnterFloatType is called when production floatType is entered.
func (s *BaseWebIDLListener) EnterFloatType(ctx *FloatTypeContext) {}

// ExitFloatType is called when production floatType is exited.
func (s *BaseWebIDLListener) ExitFloatType(ctx *FloatTypeContext) {}

// EnterUnsignedIntegerType is called when production unsignedIntegerType is entered.
func (s *BaseWebIDLListener) EnterUnsignedIntegerType(ctx *UnsignedIntegerTypeContext) {}

// ExitUnsignedIntegerType is called when production unsignedIntegerType is exited.
func (s *BaseWebIDLListener) ExitUnsignedIntegerType(ctx *UnsignedIntegerTypeContext) {}

// EnterIntegerType is called when production integerType is entered.
func (s *BaseWebIDLListener) EnterIntegerType(ctx *IntegerTypeContext) {}

// ExitIntegerType is called when production integerType is exited.
func (s *BaseWebIDLListener) ExitIntegerType(ctx *IntegerTypeContext) {}

// EnterOptionalLong is called when production optionalLong is entered.
func (s *BaseWebIDLListener) EnterOptionalLong(ctx *OptionalLongContext) {}

// ExitOptionalLong is called when production optionalLong is exited.
func (s *BaseWebIDLListener) ExitOptionalLong(ctx *OptionalLongContext) {}

// EnterPromiseType is called when production promiseType is entered.
func (s *BaseWebIDLListener) EnterPromiseType(ctx *PromiseTypeContext) {}

// ExitPromiseType is called when production promiseType is exited.
func (s *BaseWebIDLListener) ExitPromiseType(ctx *PromiseTypeContext) {}

// EnterNull_ is called when production null_ is entered.
func (s *BaseWebIDLListener) EnterNull_(ctx *Null_Context) {}

// ExitNull_ is called when production null_ is exited.
func (s *BaseWebIDLListener) ExitNull_(ctx *Null_Context) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseWebIDLListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseWebIDLListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseWebIDLListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseWebIDLListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifiers is called when production identifiers is entered.
func (s *BaseWebIDLListener) EnterIdentifiers(ctx *IdentifiersContext) {}

// ExitIdentifiers is called when production identifiers is exited.
func (s *BaseWebIDLListener) ExitIdentifiers(ctx *IdentifiersContext) {}

// EnterExtendedAttributeNoArgs is called when production extendedAttributeNoArgs is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeNoArgs(ctx *ExtendedAttributeNoArgsContext) {}

// ExitExtendedAttributeNoArgs is called when production extendedAttributeNoArgs is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeNoArgs(ctx *ExtendedAttributeNoArgsContext) {}

// EnterExtendedAttributeArgList is called when production extendedAttributeArgList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeArgList(ctx *ExtendedAttributeArgListContext) {}

// ExitExtendedAttributeArgList is called when production extendedAttributeArgList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeArgList(ctx *ExtendedAttributeArgListContext) {}

// EnterExtendedAttributeIdent is called when production extendedAttributeIdent is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeIdent(ctx *ExtendedAttributeIdentContext) {}

// ExitExtendedAttributeIdent is called when production extendedAttributeIdent is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeIdent(ctx *ExtendedAttributeIdentContext) {}

// EnterExtendedAttributeIdentList is called when production extendedAttributeIdentList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeIdentList(ctx *ExtendedAttributeIdentListContext) {}

// ExitExtendedAttributeIdentList is called when production extendedAttributeIdentList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeIdentList(ctx *ExtendedAttributeIdentListContext) {}

// EnterExtendedAttributeNamedArgList is called when production extendedAttributeNamedArgList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeNamedArgList(ctx *ExtendedAttributeNamedArgListContext) {
}

// ExitExtendedAttributeNamedArgList is called when production extendedAttributeNamedArgList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeNamedArgList(ctx *ExtendedAttributeNamedArgListContext) {
}
